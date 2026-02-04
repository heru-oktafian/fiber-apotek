package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/fiber-apotek/controllers"
	"github.com/heru-oktafian/fiber-apotek/middlewares"
)

func SysBranchRoutes(app *fiber.App) {
	// Endpoint Cabang
	app.Get("/api/branches", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), controllers.GetAllBranch)
	app.Get("/api/branches/:id", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), controllers.GetBranch)
	app.Post("/api/branches", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), controllers.CreateBranch)
	app.Put("/api/branches/:id", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), controllers.UpdateBranch)
	app.Delete("/api/branches/:id", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), controllers.DeleteBranch)
}
