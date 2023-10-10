package route

import (
	"gogod/config"
	"gogod/delivery"
	"gogod/pkg/database"
	"gogod/repository"
	"gogod/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router) {
	auth := router.Group("/auth")

	userCol := database.GetCollection(config.ENV, database.MC, "users")
	authRepo := repository.NewAuthRepository(userCol)
	userRepo := repository.NewUserRepository(userCol)

	authUcase := usecase.NewAuthUsecase(authRepo, userRepo)
	authHandler := delivery.NewAuthHandler(authUcase)

	auth.Get("/login", authHandler.Login)
}
