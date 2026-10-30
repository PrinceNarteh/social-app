package middleware

import (
	"Server/internal/config"
	"Server/internal/utils"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

func AuthGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.Env.JWT.Secret},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid or missing token")
		},
	})
}
