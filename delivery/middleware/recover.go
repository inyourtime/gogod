package middleware

import (
	"gogod/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// Recover returns a fiber.Handler function that recovers from panics and handles errors gracefully.
//
// It logs the error and responds with a 500 status code to the client if a panic occurs.
// The function takes a *fiber.Ctx parameter and returns an error.
func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Log the error
				logger.Error(r)
				// Respond with a 500 status code to the client
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"message": fiber.ErrInternalServerError.Message,
				})
			}
		}()
		// Next is called to execute the actual route handler
		return c.Next()
	}
}
