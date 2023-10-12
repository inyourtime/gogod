package delivery

import (
	"gogod/domain"
	"gogod/model"

	"github.com/go-playground/validator/v10"
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

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	if user_id == "" {
		return FiberError(c, fiber.ErrBadRequest)
	}

	req := new(model.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, err)
	}

	err := h.userUsecase.UpdateUser(user_id, req)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "update user success",
	})
}
