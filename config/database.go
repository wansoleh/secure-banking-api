package config

import (
	"fmt"
	"log"
	"secure-banking-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBInstance adalah koneksi database global
var DBInstance *gorm.DB

// InitDatabase menghubungkan ke PostgreSQL
func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Config.DBHost, Config.DBUser, Config.DBPassword, Config.DBName, Config.DBPort,
	)

	var err error
	DBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database connected successfully")
	DBInstance.AutoMigrate(&models.User{}, &models.Transaction{})
}
