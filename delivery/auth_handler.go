package delivery

import (
	"context"
	"encoding/json"
	"gogod/config"
	"gogod/domain"
	"gogod/model"
	"io"
	"net/http"

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

func (h *authHandler) Refresh(c *fiber.Ctx) error {
	type Refresh struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	req := new(Refresh)
	if err := c.BodyParser(req); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validator.New().Struct(req); err != nil {
		return FiberError(c, err)
	}

	resp, err := h.authUsecase.Refresh(req.RefreshToken)
	if err != nil {
		return FiberError(c, err)
	}
	return c.JSON(resp)
}

func (h *authHandler) Google(c *fiber.Ctx) error {
	url := config.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func (h *authHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googleCon := config.GoogleLoginConfig

	token, err := googleCon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	info := model.GoogleInfo{}
	if err := json.Unmarshal(userData, &info); err != nil {
		return c.SendString("JSON Unmarshal Failed")
	}

	userResp, err := h.authUsecase.Google(info)
	if err != nil {
		return FiberError(c, err)
	}

	return c.JSON(userResp)
}
