package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/fiber-apotek/configs"
	"github.com/heru-oktafian/fiber-apotek/routers"
	"github.com/heru-oktafian/fiber-apotek/seeders"
	"github.com/heru-oktafian/fiber-apotek/services"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// Init global timezone
	configs.InitTimezone()
	log.Println("🕒 Sekarang WIB:", time.Now().In(configs.Location))

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment
	serverPort := os.Getenv("SERVER_PORT")

	// Initialize database
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
	app := fiber.New()

	// Setup routes
	routers.AuthRoutes(app)

	// Start Fiber server
	log.Fatal(app.Listen(":" + serverPort))
}
