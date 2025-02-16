package routes

import (
	"github.com/gofiber/fiber/v2"
	"secure-banking-api/handlers"
	"secure-banking-api/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	userRoutes := api.Group("/users")
	userRoutes.Post("/register", handlers.RegisterUserHandler)
	userRoutes.Get("/:account_number", handlers.GetUserHandler)
	userRoutes.Get("/:account_number/balance", handlers.GetBalanceHandler)

	transactionRoutes := api.Group("/transactions")
	transactionRoutes.Use(middleware.JWTMiddleware)
	transactionRoutes.Post("/deposit", handlers.DepositHandler)
	transactionRoutes.Post("/withdraw", handlers.WithdrawHandler)
	transactionRoutes.Get("/:account_number", handlers.GetTransactionsHandler)
}
