package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig menyimpan semua konfigurasi
type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	RateLimit  int
	APIHost    string
	APIPort    string
	LogLevel   string
}

// Config adalah instance global konfigurasi
var Config AppConfig

// LoadConfig membaca konfigurasi dari environment variables & argument parser
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file tidak ditemukan, menggunakan environment variables sistem...")
	}

	Config.DBHost = getEnv("DB_HOST", "localhost")
	Config.DBPort = getEnv("DB_PORT", "5432")
	Config.DBUser = getEnv("DB_USER", "postgres")
	Config.DBPassword = getEnv("DB_PASSWORD", "password")
	Config.DBName = getEnv("DB_NAME", "mydb")
	Config.JWTSecret = getEnv("JWT_SECRET", "default-secret")

	rateLimit, err := strconv.Atoi(getEnv("RATE_LIMIT_MAX_REQUESTS", "100"))
	if err != nil {
		rateLimit = 100
	}
	Config.RateLimit = rateLimit

	flag.StringVar(&Config.APIHost, "api-host", "0.0.0.0", "REST API host")
	flag.StringVar(&Config.APIPort, "api-port", "8080", "REST API port")
	flag.StringVar(&Config.LogLevel, "log-level", "info", "Logging level (info, warning, error)")

	flag.Parse()
	fmt.Println("✅ Konfigurasi aplikasi berhasil dimuat!")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
