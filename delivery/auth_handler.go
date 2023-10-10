package delivery

import (
	"gogod/domain"

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
	return c.JSON("login")
}
