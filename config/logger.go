package config

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger digunakan untuk mencatat aktivitas aplikasi
var Logger *logrus.Logger

// InitLogger mengatur konfigurasi log
func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	file, err := os.OpenFile("logs/banking.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.SetOutput(os.Stdout)
	}

	switch Config.LogLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}
