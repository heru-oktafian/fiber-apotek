package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
	export_services "github.com/heru-oktafian/fiber-apotek/services/exports"
	excels "github.com/heru-oktafian/fiber-apotek/services/exports/excel/masters"
)

func ExportExcelRoutes(app *fiber.App) {
	// Inisialisasi Excel Service dan Product Handler
	excelService := export_services.NewExcelService(configs.DB)
	productHandler := excels.NewProductHandler(excelService)

	// Protected route (asumsikan sudah ada JWT middleware)
	app.Get("/api/products/export", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), productHandler.ExportExcel)
}
