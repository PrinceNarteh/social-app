package handlers

import (
	"Server/internal/models"
	"Server/internal/services"
	"Server/internal/utils"

	"github.com/gofiber/fiber/v3"
)

var _ IAuthHandler = (*AuthHandler)(nil)

type IAuthHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}

type AuthHandler struct {
	svc *services.Services
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.Bind().Body(loginRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := loginRequest.Validate(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, tokenPair, err := h.svc.Auth.Login(loginRequest)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/api/auth/refresh",
		Expires:  tokenPair.RefreshTokenTTL,
	})

	return utils.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"user":        user,
		"accessToken": tokenPair.AccessToken,
	})
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := user.Validate(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	tokenPair, err := h.svc.Auth.Register(user)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/api/auth/refresh",
		Expires:  tokenPair.RefreshTokenTTL,
	})

	return utils.SuccessResponse(c, fiber.StatusCreated, fiber.Map{
		"user":        user,
		"accessToken": tokenPair.AccessToken,
	})
}
