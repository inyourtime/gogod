package route

import (
	"gogod/delivery"
	"gogod/delivery/middleware"
	"gogod/pkg/database"
	"gogod/repository"
	"gogod/usecase"

	"github.com/gofiber/fiber/v2"
)

// UserRoute handles the user routes in the router.
//
// It takes a fiber.Router as a parameter.
// It does not return any value.
func userRoute(router fiber.Router) {
	usr := router.Group("/user")

	userRepo := repository.NewUserRepository(database.MC)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewUserHandler(userUsecase)

	{
		usr.Get("/", middleware.JwtGuard(), userHandler.GetAll)
		usr.Get("/profile", middleware.JwtGuard(), userHandler.GetProfile)
	}
}
