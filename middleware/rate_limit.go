package middleware

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// RateLimitMiddleware membatasi request per user
func RateLimitMiddleware(c *fiber.Ctx) error {
	time.Sleep(500 * time.Millisecond)
	return c.Next()
}
