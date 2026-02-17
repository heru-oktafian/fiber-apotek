package routes

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
	export_services "github.com/heru-oktafian/fiber-apotek/services/exports"
	pdf_audits "github.com/heru-oktafian/fiber-apotek/services/exports/pdf/audits"
	pdf_masters "github.com/heru-oktafian/fiber-apotek/services/exports/pdf/masters"
)

func ExportPDFRoutes(app *fiber.App) {
	log.Println("[Route] Initializing ExportPDFRoutes")

	// Inisialisasi PDF Service
	pdfService := export_services.NewPDFService(configs.DB)

	// Inisialisasi handlers untuk semua entitas
	pdfProductHandler := pdf_masters.NewPdfProductHandler(pdfService)
	pdfUnitHandler := pdf_masters.NewPdfUnitHandler(pdfService)
	pdfProductCategoryHandler := pdf_masters.NewPdfProductCategoryHandler(pdfService)
	pdfUnitConversionHandler := pdf_masters.NewPdfUnitConversionHandler(pdfService)
	pdfSupplierHandler := pdf_masters.NewPdfSupplierHandler(pdfService)
	pdfSupplierCategoryHandler := pdf_masters.NewPdfSupplierCategoryHandler(pdfService)
	pdfMemberCategoryHandler := pdf_masters.NewPdfMemberCategoryHandler(pdfService)
	pdfMemberHandler := pdf_masters.NewPdfMemberHandler(pdfService)
	pdfFirstStockHandler := pdf_audits.NewPdfFirstStockHandler(pdfService)
	pdfOpnameHandler := pdf_audits.NewPdfOpnameHandler(pdfService)

	// Protected routes untuk export PDF
	app.Get("/api/products/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfProductHandler.ExportPDF)
	app.Get("/api/units/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfUnitHandler.ExportPDF)
	app.Get("/api/product-categories/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfProductCategoryHandler.ExportPDF)
	app.Get("/api/unit-conversions/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfUnitConversionHandler.ExportPDF)
	app.Get("/api/suppliers/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfSupplierHandler.ExportPDF)
	app.Get("/api/supplier-categories/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfSupplierCategoryHandler.ExportPDF)
	app.Get("/api/member-categories/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfMemberCategoryHandler.ExportPDF)
	app.Get("/api/members/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfMemberHandler.ExportPDF)
	app.Get("/api/first-stocks/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfFirstStockHandler.ExportPDF)
	app.Get("/api/opnames/pdf", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), pdfOpnameHandler.ExportPDF)

	log.Println("[Route] ExportPDFRoutes initialized successfully!")
}
