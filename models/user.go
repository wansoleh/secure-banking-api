package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName      string        `json:"full_name"`
	NIK           string        `json:"nik" gorm:"unique"`
	PhoneNumber   string        `json:"phone_number" gorm:"unique"`
	AccountNumber string        `json:"account_number" gorm:"unique"`
	Balance       int           `json:"balance" gorm:"default:0"`
	Transactions  []Transaction `gorm:"foreignKey:AccountNumber;references:AccountNumber"`
}
