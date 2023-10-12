package domain

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrUserNotFound      = fiber.NewError(fiber.StatusNotFound, "User not found")
	ErrInvalidObjID      = fiber.NewError(fiber.StatusBadRequest, "Invalid object id")
	ErrEmailPwdIncorrect = fiber.NewError(fiber.StatusUnauthorized, "Email or Password are not correct 🥲")
	ErrEmailExist        = fiber.NewError(fiber.StatusUnprocessableEntity, "Email already exist 😜")
)
