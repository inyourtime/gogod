package middleware

import (
	"gogod/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// AuthGuard is a function that returns a fiber.Handler.
//
// It creates a new jwtware with the provided configuration and returns a fiber.Handler.
// The jwtware is configured with a signing key obtained from the Jwt.Secret configuration.
// The SuccessHandler is a function that calls the Next() method on the fiber.Ctx object.
// The ErrorHandler is a function that returns a JSON response with a status code of fiber.StatusUnauthorized,
// and includes the error message and description in the response body.
// The function returns the configured jwtware as a fiber.Handler.
func AuthGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.ENV.Jwt.Secret)},
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":        fiber.StatusUnauthorized,
				"message":     fiber.ErrUnauthorized.Message,
				"description": err.Error(),
			})
		},
	})
}
