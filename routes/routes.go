package routes

import (
	"github.com/gofiber/fiber/v2"
	"secure-banking-api/handlers"
	"secure-banking-api/middleware"
)

// SetupRoutes mengatur rute API
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Rute pengguna (nasabah)
	userRoutes := api.Group("/nasabah")
	userRoutes.Post("/daftar", handlers.RegisterUserHandler)  // Registrasi nasabah baru
	userRoutes.Get("/saldo/:no_rekening", handlers.GetBalanceHandler) // Cek saldo

	// Rute transaksi
	transactionRoutes := api.Group("/transaksi")
	transactionRoutes.Use(middleware.JWTMiddleware)
	transactionRoutes.Post("/tabung", handlers.DepositHandler)   // Menabung
	transactionRoutes.Post("/tarik", handlers.WithdrawHandler)   // Menarik dana
}
