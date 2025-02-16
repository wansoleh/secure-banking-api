package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"secure-banking-api/config"
	"secure-banking-api/models"
	"secure-banking-api/services" // Import the services package
)

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		config.Log.WithError(err).Error("Failed to parse user registration request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	// Check if NIK or PhoneNumber already exists
	var existingUser models.User
	if err := config.DBInstance.Where("nik = ? OR phone_number = ?", user.NIK, user.PhoneNumber).First(&existingUser).Error; err == nil {
		config.Log.WithFields(logrus.Fields{
			"nik":         user.NIK,
			"phoneNumber": user.PhoneNumber,
		}).Warning("User registration failed: Duplicate NIK or PhoneNumber")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "NIK atau Nomor HP sudah digunakan"})
	}

	// Generate unique account number
	user.AccountNumber = services.GenerateAccountNumber()
	user.Balance = 0

	if err := config.DBInstance.Create(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to register new user")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Gagal mendaftarkan pengguna"})
	}

	config.Log.WithFields(logrus.Fields{
		"fullName":      user.FullName,
		"nik":           user.NIK,
		"phoneNumber":   user.PhoneNumber,
		"accountNumber": user.AccountNumber,
	}).Info("New user registered successfully")

	return c.JSON(fiber.Map{"account_number": user.AccountNumber})
}
