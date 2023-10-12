package middleware

import (
	"gogod/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthGuard is a function that returns a fiber.Handler.
//
// It creates a new jwtware with the provided configuration and returns a fiber.Handler.
// The jwtware is configured with a signing key obtained from the Jwt.Secret configuration.
// The SuccessHandler is a function that calls the Next() method on the fiber.Ctx object.
// The ErrorHandler is a function that returns a JSON response with a status code of fiber.StatusUnauthorized,
// and includes the error message and description in the response body.
// The function returns the configured jwtware as a fiber.Handler.
func JwtGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.ENV.Jwt.Secret)},
		SuccessHandler: func(c *fiber.Ctx) error {
			if claims, ok := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims); ok {
				c.Locals("user_id", claims["user_id"])
				c.Locals("email", claims["email"])
			}
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
