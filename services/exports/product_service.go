package services

import (
	"fmt"
	"strconv"

	"github.com/heru-oktafian/fiber-apotek/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExcelService struct {
	db *gorm.DB
}

func NewExcelService(db *gorm.DB) *ExcelService {
	return &ExcelService{db: db}
}

// ExportProductsToExcel menghasilkan file Excel sesuai format yang kamu minta
func (s *ExcelService) ExportProductsToExcel(branchID string) ([]byte, error) {
	var products []models.ProductDetail

	// Query dengan join agar dapat UnitName (kalau ProductDetail kamu sudah include, bisa di-skip join-nya)
	err := s.db.Debug().Model(&models.Product{}).
		Select(`
			products.*,
			units.name as unit_name,
			product_categories.name as product_category_name
		`).
		Joins("LEFT JOIN units ON units.id = products.unit_id").
		Joins("LEFT JOIN product_categories ON product_categories.id = products.product_category_id").
		//Where("products.branch_id = ?", branchID).
		Order("products.name ASC").
		Find(&products).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	f := excelize.NewFile()
	sheet := "Produk"
	f.SetSheetName("Sheet1", sheet)

	// Header
	headers := []string{
		"ID", "SKU", "NAME", "ALIAS",
		"PURCHASE PRI", "SALE PRI", "ALTERNATIF PRI",
		"STOCK", "UNIT", "EXPIRED DATE",
	}

	for i, h := range headers {
		cell, _ := excelize.ColumnNumberToName(i + 1)
		f.SetCellValue(sheet, cell+"1", h)
	}

	// Style Header (bold + background)
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#1E88E5"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	f.SetCellStyle(sheet, "A1", "J1", headerStyle)

	// Isi Data
	for i, p := range products {
		row := i + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), p.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), p.SKU)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), p.Name)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), p.Alias)

		// Harga dengan format Rp. x.xxx
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), formatRupiah(p.PurchasePrice))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), formatRupiah(p.SalesPrice))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), formatRupiah(p.AlternatePrice))

		// STOCK + UNIT digabung seperti di screenshot kamu (500 PCS)
		stockUnit := fmt.Sprintf("%d %s", p.Stock, p.UnitName)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), stockUnit)

		// UNIT (nama satuan saja)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), p.UnitName)

		// Expired Date → DD/MM/YYYY
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), p.ExpiredDate.Format("02/01/2006"))
	}

	// Auto width kolom
	_ = f.SetColWidth(sheet, "A", "A", 18)
	_ = f.SetColWidth(sheet, "B", "B", 15)
	_ = f.SetColWidth(sheet, "C", "C", 35)
	_ = f.SetColWidth(sheet, "D", "D", 25)
	_ = f.SetColWidth(sheet, "E", "G", 15)
	_ = f.SetColWidth(sheet, "H", "H", 12)
	_ = f.SetColWidth(sheet, "I", "I", 10)
	_ = f.SetColWidth(sheet, "J", "J", 15)

	// Buat Table (supaya ada filter & styling bagus)
	_ = f.AddTable(sheet, &excelize.Table{
		Range:           fmt.Sprintf("A1:J%d", len(products)+1),
		Name:            "ProdukTable",
		StyleName:       "TableStyleMedium9",
		ShowFirstColumn: false,
		ShowLastColumn:  false,
		// ShowRowStripes:    true,
		ShowColumnStripes: false,
	})

	// Return sebagai byte (siap dikirim ke browser)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), nil
}

// formatRupiah mengubah 4900 → "Rp. 4.900"
func formatRupiah(amount int) string {
	if amount == 0 {
		return "Rp. 0"
	}
	s := strconv.Itoa(amount)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "." + s[i:]
	}
	return "Rp. " + s
}
