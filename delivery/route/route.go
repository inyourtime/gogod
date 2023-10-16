package route

import "github.com/gofiber/fiber/v2"

// SetupRoute sets up the routes for the application.
//
// Parameters:
//   - app: a pointer to a fiber.App instance
//
// Return type: None
func SetupRoute(app *fiber.App) {
	// healthcheck
	app.Get("/hc", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is still OK 😛",
		})
	})

	api := app.Group("/api")

	// register here
	{
		authRoute(api)
		userRoute(api)
		blogRoute(api)
	}
}
