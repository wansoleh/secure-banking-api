package services

import (
	"errors"
	"fmt"
	"math/rand"
	"secure-banking-api/config"
	"secure-banking-api/models"
)

// RegisterNewUser mendaftarkan user dengan transaksi database
func RegisterNewUser(fullName, nik, phoneNumber string) (string, error) {
	// Cek apakah NIK atau Nomor HP sudah terdaftar
	var existingUser models.User
	if err := config.DBInstance.Where("nik = ? OR phone_number = ?", nik, phoneNumber).First(&existingUser).Error; err == nil {
		return "", errors.New("NIK atau Nomor HP sudah terdaftar")
	}

	// Generate nomor rekening unik
	accountNumber := generateAccountNumber()

	// Buat user baru
	newUser := models.User{
		FullName:      fullName,
		NIK:           nik,
		PhoneNumber:   phoneNumber,
		AccountNumber: accountNumber,
		Balance:       0,
	}

	// Mulai transaksi database
	tx := config.DBInstance.Begin()
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return accountNumber, nil
}

// generateAccountNumber menghasilkan nomor rekening unik secara acak
func generateAccountNumber() string {
	return fmt.Sprintf("112233%d", rand.Intn(89999)+10000) // Nomor rekening 112233XXXXX
}
