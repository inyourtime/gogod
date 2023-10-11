package route

import (
	"gogod/config"
	"gogod/delivery"
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

	userCol := database.GetCollection(config.ENV, database.MC, "users")
	authRepo := repository.NewAuthRepository(userCol)
	userRepo := repository.NewUserRepository(userCol)

	authUcase := usecase.NewAuthUsecase(authRepo, userRepo)
	authHandler := delivery.NewAuthHandler(authUcase)

	auth.Post("/login", authHandler.Login)
}
