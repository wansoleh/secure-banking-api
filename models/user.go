package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	FullName      string        `json:"full_name" gorm:"not null"`
	NIK           string        `json:"nik" gorm:"unique;not null"`
	PhoneNumber   string        `json:"phone_number" gorm:"unique;not null"`
	AccountNumber string        `json:"account_number" gorm:"unique;not null"`
	Balance       int           `json:"balance" gorm:"default:0;not null"`
	Transactions  []Transaction `gorm:"foreignKey:UserID"` // ✅ Perbaikan foreign key
}
