package handlers

import (
	"net/http"
	"secure-banking-api/config"
	"secure-banking-api/services"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserHandler menangani pendaftaran user
func RegisterUserHandler(c *fiber.Ctx) error {
	var request struct {
		FullName    string `json:"full_name"`
		NIK         string `json:"nik"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	accountNumber, err := services.RegisterNewUser(request.FullName, request.NIK, request.PhoneNumber)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"account_number": accountNumber})
}
