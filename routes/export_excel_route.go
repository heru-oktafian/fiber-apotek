package routes

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
	export_services "github.com/heru-oktafian/fiber-apotek/services/exports"
	excels "github.com/heru-oktafian/fiber-apotek/services/exports/excel/masters"
)

func ExportExcelRoutes(app *fiber.App) {
	log.Println("[Route] Initializing ExportExcelRoutes")

	// Inisialisasi Excel Service dan Product Handler
	excelService := export_services.NewExcelService(configs.DB)
	productHandler := excels.NewProductHandler(excelService)

	// Protected route: Export produk ke Excel
	app.Get("/api/products/export", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), productHandler.ExportExcel)

	log.Println("[Route] ExportExcelRoutes initialized successfully!")
}
