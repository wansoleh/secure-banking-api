package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"secure-banking-api/config"
	"secure-banking-api/routes"
)

func main() {
	// 🔧 Load konfigurasi dari environment variables & argument parser
	config.LoadConfig()

	// 🔧 Inisialisasi database
	config.InitDatabase()

	// 🔧 Inisialisasi Logger
	config.InitLogger()

	// 🔧 Inisialisasi Fiber API
	app := fiber.New()

	// 🔧 Setup Routing
	routes.SetupRoutes(app)

	// 🔧 Jalankan server dengan API_HOST & API_PORT dari konfigurasi
	apiAddr := fmt.Sprintf("%s:%s", config.Config.APIHost, config.Config.APIPort)
	log.Printf("🚀 Secure Banking API running on %s", apiAddr)
	app.Listen(apiAddr)
}
