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

	// Debug: Cek ukuran bytes yang dikembalikan
	log.Printf("DEBUG: Ukuran excel bytes: %d bytes", len(excelBytes))

	// Cek apakah bytes kosong (kalau gak ada data)
	if len(excelBytes) == 0 {
		log.Println("ERROR: Excel bytes kosong - mungkin ada issue di service")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "gagal generate file excel - bytes kosong",
		})
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", `attachment; filename="produk.xlsx"`)
	return c.Send(excelBytes)
}
