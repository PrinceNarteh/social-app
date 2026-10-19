// Package utils
package utils

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if fiberErr, ok := errors.AsType[*fiber.Error](err); ok {
		code = fiberErr.Code
	}

	return ErrorResponse(ctx, code, err.Error())
}
