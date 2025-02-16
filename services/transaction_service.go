package services

import (
	"errors"
	"secure-banking-api/config"
	"secure-banking-api/models"
)

// Deposit menambahkan saldo
func Deposit(accountNumber string, amount int) (int, error) {
	var user models.User
	tx := config.DBInstance.Begin()

	if err := tx.Where("account_number = ?", accountNumber).First(&user).Error; err != nil {
		tx.Rollback()
		return 0, errors.New("Nomor rekening tidak ditemukan")
	}

	user.Balance += amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	transaction := models.Transaction{
		AccountNumber:  accountNumber,
		TransactionType: "deposit",
		Amount:         amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return user.Balance, nil
}
