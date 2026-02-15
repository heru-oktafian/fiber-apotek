package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
	excels "github.com/heru-oktafian/fiber-apotek/services/exports/excel/masters"
)

func SetupRoutes(app *fiber.App, handler *excels.ProductHandler) {
	// Protected route (asumsikan sudah ada JWT middleware)
	app.Get("/api/products/export", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), handler.ExportExcel)
}
