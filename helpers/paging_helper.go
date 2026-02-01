package helpers

import (
	"math"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	models "github.com/heru-oktafian/fiber-apotek/models"
	"gorm.io/gorm"
)

// Paginate adalah helper untuk menangani pagination dan search pada query
func Paginate(c *fiber.Ctx, query *gorm.DB, model interface{}, searchFields []string) (interface{}, string, int, int, int, int, error) {
	// Parsing body JSON ke struct
	var body models.RequestBody
	if err := c.BodyParser(&body); err != nil {
		return nil, "", 0, 0, 0, 0, err
	}

	// Validasi dan set default untuk page jika tidak valid
	page := body.Page
	if page < 1 {
		page = 1
	}
	limit := 12                              // Tetapkan limit ke 12 data per halaman
	search := strings.TrimSpace(body.Search) // Ambil search key dari body
	offset := (page - 1) * limit

	// Jika ada search key, tambahkan filter WHERE
	if search != "" {
		search = strings.ToLower(search) // Konversi search ke lowercase
		if len(searchFields) > 0 {
			whereClause := ""
			args := make([]interface{}, len(searchFields))
			for i, field := range searchFields {
				if whereClause != "" {
					whereClause += " OR "
				}
				whereClause += "LOWER(" + field + ") LIKE ?"
				args[i] = "%" + search + "%"
			}
			query = query.Where(whereClause, args...)
		}
	}

	var total int64

	// Hitung total unit yang sesuai dengan filter
	if err := query.Count(&total).Error; err != nil {
		return nil, "", 0, 0, 0, 0, err
	}

	// Ambil data dengan pagination
	if err := query.Offset(offset).Limit(limit).Scan(model).Error; err != nil {
		return nil, "", 0, 0, 0, 0, err
	}

	// Hitung total halaman berdasarkan hasil filter
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return model, search, int(total), page, totalPages, limit, nil
}
