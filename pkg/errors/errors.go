package errors

import (
	"errors"
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

func Is(err error, target error) bool {
	ftarget, okt := target.(*fiber.Error)
	ferr, oke := err.(*fiber.Error)

	if okt && oke {
		return ftarget.Code == ferr.Code
	}

	return errors.Is(err, target)
}
