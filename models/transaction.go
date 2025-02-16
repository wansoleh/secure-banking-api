package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	AccountNumber  string `json:"account_number"`
	TransactionType string `json:"transaction_type" gorm:"type:varchar(10);check:transaction_type IN ('deposit', 'withdraw')"`
	Amount         int    `json:"amount" gorm:"check:amount > 0"`
}
