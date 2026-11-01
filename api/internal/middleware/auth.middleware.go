package middleware

import (
	"strings"

	"Server/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func AuthGuard(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header not present")
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, `Invalid token format. It should be "Bearer token"`)
	}

	claims, err := utils.VerifyToken(tokenSlice[1])
	if err != nil {
		return err
	}

	c.Locals("userId", claims.Subject)

	return c.Next()
}
