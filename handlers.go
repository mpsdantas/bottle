package bottle

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mpsdantas/bottle/pkg/log"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(&fiber.Map{
			"code":    e.Code,
			"message": e.Message,
		})
	}

	log.Error(ctx.Context(), "internal server error occurred", log.Err(err))

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"message": "internal server error",
	})
}
