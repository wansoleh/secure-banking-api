package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"secure-banking-api/config"
	"secure-banking-api/models"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		config.Log.WithError(err).Error("Failed to parse user registration request")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request"})
	}

	// Cek apakah NIK atau NoHP sudah terdaftar
	var existingUser models.User
	if err := config.DB.Where("nik = ? OR no_hp = ?", user.NIK, user.NoHP).First(&existingUser).Error; err == nil {
		config.Log.WithFields(logrus.Fields{
			"nik":   user.NIK,
			"no_hp": user.NoHP,
		}).Warning("User registration failed: Duplicate NIK or No HP")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "NIK atau No HP sudah digunakan"})
	}

	// Generate nomor rekening unik
	user.NoRekening = generateNoRekening()
	user.Saldo = 0

	if err := config.DB.Create(&user).Error; err != nil {
		config.Log.WithError(err).Error("Failed to register new user")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Gagal mendaftarkan pengguna"})
	}

	config.Log.WithFields(logrus.Fields{
		"nama":        user.Nama,
		"nik":         user.NIK,
		"no_hp":       user.NoHP,
		"no_rekening": user.NoRekening,
	}).Info("New user registered successfully")

	return c.JSON(fiber.Map{"no_rekening": user.NoRekening})
}

func generateNoRekening() string {
	return "112233" + string(rand.Intn(99999-10000)+10000)
}
