package delivery

import (
	"gogod/domain"
	"gogod/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(au domain.AuthUsecase) *authHandler {
	return &authHandler{
		authUsecase: au,
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	req := new(model.AuthLoginRequest)
	if err := c.BodyParser(req); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, err)
	}

	res, err := h.authUsecase.Login(req)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(res)
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	req := new(model.User)
	if err := c.BodyParser(req); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, err)
	}
	req.Provider = model.LocalProvider

	res, err := h.authUsecase.Register(req)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(res)
}
