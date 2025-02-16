package config

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Log digunakan untuk mencatat aktivitas aplikasi
var Log *logrus.Logger

// InitLogger mengatur konfigurasi log
func InitLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Pastikan direktori logs/ ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", 0755) // Buat folder jika belum ada
	}

	// Buka file log
	file, err := os.OpenFile("logs/banking.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
		Log.Warn("Gagal membuka file log, menggunakan stdout")
	}

	// Ambil level log dari environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}
