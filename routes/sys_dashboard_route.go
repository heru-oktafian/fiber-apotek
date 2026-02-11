package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	controllers "github.com/heru-oktafian/fiber-apotek/controllers"
	middlewares "github.com/heru-oktafian/fiber-apotek/middlewares"
)

func SysDashboardRoute(app *fiber.App) {
	app.Get("/api/sys/dashboard/monthly-profit-report", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.MonthlyProfitReport)
	app.Get("/api/sys/dashboard/weekly-profit-report", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.WeeklyProfitReport)
	app.Get("/api/sys/dashboard/daily-profit-report", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.DailyProfitReport)
	app.Get("/api/sys/dashboard/top-selling-products", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.GetTopSellingProducts)
	app.Get("/api/sys/dashboard/least-selling-products", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.GetLeastSellingProducts)
	app.Get("/api/sys/dashboard/neared-report", middlewares.JWTMiddleware, middlewares.RoleMiddleware("administrator", "operator", "cashier", "finance", "superadmin"), controllers.GetExpiringProducts)
}
