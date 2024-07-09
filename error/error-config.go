package error

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Create a new fiber instance with custom config
var ErrorConfig = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		// Send custom error page
		err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
		if err != nil {
			// In case the SendFile fails
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Return from handler
		return nil
	},
}
