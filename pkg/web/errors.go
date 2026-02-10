package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotFound(msg string) error {
	return fiber.NewError(fiber.StatusNotFound, msg)
}

func Internal(msg string) error {
	return fiber.NewError(fiber.StatusInternalServerError, msg)
}

func Validation(msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func FailureCause(cause string, msg string) error {
	return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("%v: %v", cause, msg))
}

func Unauthorized(msg string) error {
	return fiber.NewError(fiber.StatusUnauthorized, msg)
}

func Forbidden(msg string) error {
	return fiber.NewError(fiber.StatusForbidden, msg)
}

func Error(code int, text string) error {
	return fiber.NewError(code, text)
}
