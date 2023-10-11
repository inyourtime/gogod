package delivery

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type validationErrResponse struct {
	Field string `json:"field"`
	Tag   string `json:"error"`
}

func FiberError(c *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case *fiber.Error:
		return c.Status(err.Code).JSON(err)
	case validator.ValidationErrors:
		var errs []validationErrResponse
		for _, e := range err {
			errs = append(errs, validationErrResponse{Field: e.Field(), Tag: e.Tag()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":        fiber.StatusBadRequest,
			"message":     fiber.ErrBadRequest.Message,
			"description": errs,
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}
