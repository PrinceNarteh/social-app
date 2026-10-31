package utils

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success    bool `json:"success"`
	StatusCode int  `json:"statusCode"`
	Data       any  `json:"data,omitempty"`
	Error      any  `json:"error,omitempty"`
}

func SuccessResponse(ctx fiber.Ctx, statusCode int, data any) error {
	return ctx.Status(statusCode).JSON(Response{
		Success:    true,
		StatusCode: statusCode,
		Data:       data,
	})
}

func ErrorResponse(ctx fiber.Ctx, statusCode int, err any) error {
	return ctx.Status(statusCode).JSON(Response{
		Success:    false,
		StatusCode: statusCode,
		Error:      err,
	})
}
