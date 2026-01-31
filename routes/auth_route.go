package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/heru-oktafian/fiber-apotek/controllers"
	"github.com/heru-oktafian/fiber-apotek/helpers"
	"github.com/heru-oktafian/fiber-apotek/middlewares"
	"github.com/heru-oktafian/fiber-apotek/services"
)

func AuthRoutes(app *fiber.App) {
	// Adding logger middleware of fiber
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Sesuaikan jika kamu ingin membatasi domain tertentu
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/coba", func(c *fiber.Ctx) error {
		// defaultMember := helpers.GetMemberDefault(config.DB, "BRC250118132203")
		// // return defaultMember
		// return c.SendString(defaultMember)
		return c.SendString("Halo dari Fiber di port " + os.Getenv("SERVER_PORT"))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return helpers.JSONResponse(c, fiber.StatusOK, "Pesan anda telah kami terima dan segera kami tindak lanjuti.", nil)
	})

	app.Get("/files/dump", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), func(c *fiber.Ctx) error {
		files, err := services.ListDumpFiles()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(files)
	})

	app.Get("/files/rest", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), func(c *fiber.Ctx) error {
		files, err := services.ListRestFiles()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(files)
	})

	// Auth Endpoints
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	// app.Post("/register", controllers.CreateUser)
	app.Get("/api/profile", middlewares.JWTMiddleware, controllers.GetProfile)
	// app.Get("/list_branches", middlewares.JWTMiddleware, controllers.GetBranchByUserId)

	// SetBranch Endpoint
	app.Post("/api/set_branch", controllers.SetBranch)
	// api := app.Group("/api", middlewares.JWTMiddleware)

	// Endpoint to generate file .env
	app.Post("/api/update-env", middlewares.JWTMiddleware, middlewares.RoleMiddleware("superadmin", "administrator"), func(c *fiber.Ctx) error {
		type request struct {
			Content string `json:"content"`
		}

		var body request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Gagal membaca request body",
			})
		}

		if body.Content == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Content tidak boleh kosong",
			})
		}

		if err := services.WriteRawEnvFile(".env", body.Content); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "✅ File .env berhasil diperbarui",
		})
	})
}
