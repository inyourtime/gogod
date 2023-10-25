package route

import (
	"gogod/delivery"
	"gogod/delivery/middleware"
	"gogod/pkg/database"
	"gogod/repository"
	"gogod/usecase"

	"github.com/gofiber/fiber/v2"
)

// AuthRoute is a function that sets up the authentication routes for the given router.
//
// It takes a `router` parameter of type `fiber.Router` which represents the router object to set up the routes on.
// The function does not return anything.
func authRoute(router fiber.Router) {
	auth := router.Group("/auth")

	authRepo := repository.NewAuthRepository(database.MC)
	userRepo := repository.NewUserRepository(database.MC)

	authUcase := usecase.NewAuthUsecase(authRepo, userRepo)
	authHandler := delivery.NewAuthHandler(authUcase)

	{
		auth.Post("/login", authHandler.Login)
		auth.Post("/register", authHandler.Register)
		auth.Post("/refresh", authHandler.Refresh)
		auth.Get("/test", middleware.JwtGuard(), func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"user_id": c.Locals("user_id"),
				"email":   c.Locals("email"),
			})
		})
		auth.Get("/google", authHandler.Google)
		auth.Get("/google/callback", authHandler.GoogleCallback)
	}

}
