package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10, // Max 10 requests per minute
		Expiration: time.Minute,
	})
}
