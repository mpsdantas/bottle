package web

import (
	"github.com/gofiber/fiber/v2"
)

func Error(code int, text string) error {
	return fiber.NewError(code, text)
}
