package services

import (
	"errors"
	"math/rand"
	"secure-banking-api/config"
	"secure-banking-api/models"

	"gorm.io/gorm"
)

// RegisterNewUser mendaftarkan user dengan transaksi database
func RegisterNewUser(fullName, nik, phoneNumber string) (string, error) {
	var existingUser models.User
	if err := config.DBInstance.Where("nik = ? OR phone_number = ?", nik, phoneNumber).First(&existingUser).Error; err == nil {
		return "", errors.New("NIK atau Nomor HP sudah terdaftar")
	}

	accountNumber := generateAccountNumber()
	newUser := models.User{
		FullName:      fullName,
		NIK:           nik,
		PhoneNumber:   phoneNumber,
		AccountNumber: accountNumber,
		Balance:       0,
	}

	tx := config.DBInstance.Begin()
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return accountNumber, nil
}

func generateAccountNumber() string {
	return "112233" + string(rand.Intn(99999-10000)+10000)
}
