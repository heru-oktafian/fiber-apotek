package excels

import (
	fmt "fmt"

	fiber "github.com/gofiber/fiber/v2"
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
	branchID, ok := c.Locals("branch_id").(string)
	if !ok || branchID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "branch_id tidak ditemukan di context",
		})
	}

	excelBytes, err := h.excelService.ExportProductsToExcel(branchID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("gagal generate excel: %v", err),
		})
	}

	// Set header supaya langsung download
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", `attachment; filename="produk.xlsx"`)

	// Kirim file
	return c.Send(excelBytes)
}
