package excels

import (
	fmt "fmt"
	log "log"

	fiber "github.com/gofiber/fiber/v2"
	services "github.com/heru-oktafian/fiber-apotek/services"
	export_services "github.com/heru-oktafian/fiber-apotek/services/exports"
)

type ProductHandler struct {
	excelService *export_services.ExcelService
}

func NewProductHandler(excelService *export_services.ExcelService) *ProductHandler {
	return &ProductHandler{excelService: excelService}
}

// ExportExcel → Export ke Excel menggunakan Fiber
func (h *ProductHandler) ExportExcel(c *fiber.Ctx) error {
	// Ambil branch_id dari JWT middleware kamu (biasanya disimpan di Locals)
	branchID, _ := services.GetBranchID(c)
	log.Printf("DEBUG: Export dimulai untuk branch_id: %s", branchID)

	excelBytes, err := h.excelService.ExportProductsToExcel(branchID)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("gagal generate excel: %v", err),
		})
	}

	// Cek apakah bytes kosong (kalau gak ada data)
	if len(excelBytes) == 0 {
		log.Println("WARNING: Excel generated tapi kosong (no data)")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Data ditemukan tapi kosong",
			"data":    nil, // Ini yang bikin response kamu gini
		})
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", `attachment; filename="produk.xlsx"`)
	return c.Send(excelBytes)
}
