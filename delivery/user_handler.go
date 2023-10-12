package delivery

import (
	"gogod/domain"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (h *userHandler) GetProfile(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	if user_id == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}
	user, err := h.userUsecase.GetProfile(user_id)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(user)
}

func (h *userHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAllUser()
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(users)
}
