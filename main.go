package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/heru-oktafian/fiber-apotek/configs"
	"github.com/heru-oktafian/fiber-apotek/helpers"
	"github.com/heru-oktafian/fiber-apotek/routes"
	"github.com/heru-oktafian/fiber-apotek/seeders"
	"github.com/heru-oktafian/fiber-apotek/services"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// Inisialisasi zona waktu global
	configs.InitTimezone()
	log.Println("🕒 Sekarang WIB:", time.Now().In(configs.Location))

	// Muat file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Dapatkan port dari environment
	serverPort := os.Getenv("SERVER_PORT")

	// Inisialisasi database
	if err := configs.SetupDB(); err != nil {
		log.Fatal(err)
	}

	// Jalankan seeder jika ada argumen "seed"
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeders.UserSeed()
		seeders.BranchSeed()
		seeders.UserBranchSeed()
		os.Exit(0)
	}

	go func() {
		// 1. Set lokasi waktu ke Asia/Jakarta
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			log.Fatalf("❌ Gagal load lokasi Asia/Jakarta: %v", err)
		}

		// 2. Inisialisasi cron dengan lokasi waktu yang benar
		c := cron.New(cron.WithLocation(loc))

		// 3. Tambahkan job harian pukul 23:30
		_, err = c.AddFunc("30 23 * * *", func() {
			log.Println("🚀 Memulai backup database dan upload ke Google Drive...")

			if err := services.DumpDatabaseToFile(); err != nil {
				log.Printf("❌ Backup gagal: %v\n", err)
			} else {
				log.Println("✅ Backup dan upload ke Google Drive berhasil.")
			}
		})

		if err != nil {
			log.Fatalf("❌ Gagal menjadwalkan cron job: %v", err)
		}

		// 4. Mulai cron scheduler
		c.Start()
		log.Println("📅 Cron backup database aktif setiap pukul 23:30 WIB")
	}()

	// Inisialisasi Fiber
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // ⛔ wajib
	})

	// Menambahkan middleware logger fiber
	app.Use(logger.New())

	// Aktifkan CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Sesuaikan jika kamu ingin membatasi domain tertentu
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup rute
	routes.AuthRoutes(app)

	// Hitung total rute
	routeCount := 0
	for _, routes := range app.Stack() {
		routeCount += len(routes)
	}

	port, err := strconv.Atoi(serverPort)
	if err != nil {
		log.Fatal("Invalid SERVER_PORT: must be a number")
	}

	helpers.PrintFiberLikeBanner(
		os.Getenv("APPNAME"),
		"0.0.0.0",
		port,
		routeCount, // jumlah handlers
	)

	// Jalankan server Fiber
	log.Fatal(app.Listen(":" + serverPort))
}
