package models

import "gorm.io/gorm"

// Transaction model
type Transaction struct {
	gorm.Model
	UserID        uint   `json:"user_id"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Type          string `json:"type"` // ✅ Ganti TransactionType → Type
}
