package utils

import "github.com/gofiber/fiber/v3"

type response struct {
	Success    bool `json:"success"`
	StatusCode int  `json:"statusCode"`
	Data       any  `json:"data,omitempty"`
	Error      any  `json:"error,omitempty"`
}

func SuccessResponse(ctx fiber.Ctx, statusCode int, data any) error {
	return ctx.Status(statusCode).JSON(response{
		Success:    true,
		StatusCode: statusCode,
		Data:       data,
	})
}

func ErrorResponse(ctx fiber.Ctx, statusCode int, err any) error {
	return ctx.Status(statusCode).JSON(response{
		Success:    false,
		StatusCode: statusCode,
		Data:       err,
	})
}
