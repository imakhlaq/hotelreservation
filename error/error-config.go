package error

import (
	"github.com/gofiber/fiber/v2"
)

// Config Create a new fiber instance with custom config
var Config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]any{"status": 400, "message": err.Error()})
	},
}
