package api

import (
	"github.com/gofiber/fiber/v2"
)

func ApiMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
