package handlers

import (
	"net/http"

	"secure-banking-api/config"
	"secure-banking-api/models"
	"secure-banking-api/services" // Import the services package

	"github.com/gofiber/fiber/v2"
)

// RegisterUserHandler menangani registrasi nasabah baru
func RegisterUserHandler(c *fiber.Ctx) error {
	type Request struct {
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		config.Log.WithError(err).Error("Gagal memproses permintaan registrasi")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "Permintaan tidak valid"})
	}

	// Periksa apakah NIK atau NoHP sudah digunakan
	var existingUser models.User
	if err := config.DBInstance.Where("nik = ? OR phone_number = ?", req.NIK, req.NoHP).First(&existingUser).Error; err == nil {
		config.Log.Warn("Registrasi gagal: NIK atau No HP sudah digunakan")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"remark": "NIK atau Nomor HP sudah digunakan"})
	}

	// Generate unique account number
	noRekening := services.GenerateAccountNumber()
	newUser := models.User{
		FullName:      req.Nama,
		NIK:           req.NIK,
		PhoneNumber:   req.NoHP,
		AccountNumber: noRekening,
		Balance:       0,
	}

	// Simpan user baru ke database
	if err := config.DBInstance.Create(&newUser).Error; err != nil {
		config.Log.WithError(err).Error("Gagal menyimpan pengguna baru")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"remark": "Gagal mendaftarkan pengguna"})
	}

	config.Log.WithFields(map[string]interface{}{
		"nama":        newUser.FullName,
		"nik":         newUser.NIK,
		"no_hp":       newUser.PhoneNumber,
		"no_rekening": newUser.AccountNumber,
	}).Info("Pengguna baru berhasil terdaftar")

	return c.JSON(fiber.Map{"no_rekening": noRekening})
}

// GetBalanceHandler mengambil saldo pengguna berdasarkan nomor rekening
func GetBalanceHandler(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")

	var user models.User
	if err := config.DBInstance.Where("account_number = ?", noRekening).First(&user).Error; err != nil {
		config.Log.Warn("Gagal mengambil saldo: Nomor rekening tidak ditemukan", noRekening)
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"saldo": user.Balance})
}
