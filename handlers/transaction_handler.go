package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"secure-banking-api/config"
	"secure-banking-api/models"
)

// DepositHandler handles deposits
func DepositHandler(c *fiber.Ctx) error {
	type Request struct {
		AccountNumber string `json:"account_number"`
		Amount        int    `json:"amount"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		config.Log.WithError(err).Error("Failed to parse deposit request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	var user models.User
	if err := config.DBInstance.Where("account_number = ?", req.AccountNumber).First(&user).Error; err != nil {
		config.Log.WithFields(logrus.Fields{"account_number": req.AccountNumber}).Warning("Deposit failed: Account not found")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Account number not found"})
	}

	user.Balance += req.Amount
	if err := config.DBInstance.Save(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to update balance after deposit")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Failed to process transaction"})
	}

	config.Log.WithFields(logrus.Fields{
		"account_number": req.AccountNumber,
		"amount":         req.Amount,
		"balance":        user.Balance,
	}).Info("Deposit successful")

	return c.JSON(fiber.Map{"balance": user.Balance})
}

// WithdrawHandler handles withdrawals
func WithdrawHandler(c *fiber.Ctx) error {
	type Request struct {
		AccountNumber string `json:"account_number"`
		Amount        int    `json:"amount"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		config.Log.WithError(err).Error("Failed to parse withdrawal request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	var user models.User
	if err := config.DBInstance.Where("account_number = ?", req.AccountNumber).First(&user).Error; err != nil {
		config.Log.WithFields(logrus.Fields{"account_number": req.AccountNumber}).Warning("Withdrawal failed: Account not found")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Account number not found"})
	}

	if user.Balance < req.Amount {
		config.Log.WithFields(logrus.Fields{
			"account_number": req.AccountNumber,
			"amount":         req.Amount,
			"balance":        user.Balance,
		}).Warning("Withdrawal failed: Insufficient balance")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Insufficient balance"})
	}

	user.Balance -= req.Amount
	if err := config.DBInstance.Save(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to update balance after withdrawal")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Failed to process transaction"})
	}

	config.Log.WithFields(logrus.Fields{
		"account_number": req.AccountNumber,
		"amount":         req.Amount,
		"balance":        user.Balance,
	}).Info("Withdrawal successful")

	return c.JSON(fiber.Map{"balance": user.Balance})
}

// GetTransactionsHandler retrieves transactions by account number
func GetTransactionsHandler(c *fiber.Ctx) error {
	accountNumber := c.Params("account_number")

	var transactions []models.Transaction
	if err := config.DBInstance.Where("account_number = ?", accountNumber).Find(&transactions).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"remark": "Transactions not found"})
	}

	return c.JSON(transactions)
}
