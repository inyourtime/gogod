package route

import (
	"gogod/config"
	"gogod/delivery"
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

	userCol := database.GetCollection(config.ENV, database.MC, "users")
	userRepo := repository.NewUserRepository(userCol)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewUserHandler(userUsecase)

	_ = userHandler
	usr.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("user")
	})
}
