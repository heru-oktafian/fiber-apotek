package routes

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
	export_services "github.com/heru-oktafian/fiber-apotek/services/exports"
	excel_audits "github.com/heru-oktafian/fiber-apotek/services/exports/excel/audits"
	excel_masters "github.com/heru-oktafian/fiber-apotek/services/exports/excel/masters"
)

func ExportExcelRoutes(app *fiber.App) {
	log.Println("[Route] Initializing ExportExcelRoutes")

	// Inisialisasi Excel Service dan semua Handlers
	excelService := export_services.NewExcelServices(configs.DB)
	excelProductHandler := excel_masters.NewExcelProductHandler(excelService)
	excelUnitHandler := excel_masters.NewExcelUnitHandler(excelService)
	excelProductCategoryHandler := excel_masters.NewExcelProductCategoryHandler(excelService)
	excelUnitConversionHandler := excel_masters.NewExcelUnitConversionHandler(excelService)
	excelSupplierHandler := excel_masters.NewExcelSupplierHandler(excelService)
	excelSupplierCategoryHandler := excel_masters.NewExcelSupplierCategoryHandler(excelService)
	excelMemberCategoryHandler := excel_masters.NewExcelMemberCategoryHandler(excelService)
	excelMemberHandler := excel_masters.NewExcelMemberHandler(excelService)
	excelFirstStockHandler := excel_audits.NewExcelFirstStockHandler(excelService)
	excelOpnameHandler := excel_audits.NewExcelOpnameHandler(excelService)

	// Protected routes: Export ke Excel
	app.Get("/api/products/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelProductHandler.ExportExcel)
	app.Get("/api/units/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelUnitHandler.ExportExcel)
	app.Get("/api/product-categories/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelProductCategoryHandler.ExportExcel)
	app.Get("/api/unit-conversions/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelUnitConversionHandler.ExportExcel)
	app.Get("/api/suppliers/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelSupplierHandler.ExportExcel)
	app.Get("/api/supplier-categories/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelSupplierCategoryHandler.ExportExcel)
	app.Get("/api/member-categories/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelMemberCategoryHandler.ExportExcel)
	app.Get("/api/members/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelMemberHandler.ExportExcel)
	app.Get("/api/first-stocks/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelFirstStockHandler.ExportExcel)
	app.Get("/api/opnames/excel", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), excelOpnameHandler.ExportExcel)

	log.Println("[Route] ExportExcelRoutes initialized successfully!")
}
