package handlers

import (
	"Server/internal/models"
	"Server/internal/services"
	"Server/internal/utils"
	"fmt"

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

// Register
// @Summary Register a new user
// @Description Register a new user by providing firstName, lastName, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "User register details"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]any
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	registerRequest := new(models.RegisterRequest)
	if err := c.Bind().Body(registerRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	fmt.Println(registerRequest)
	if err := registerRequest.Validate(); err != nil {
		fmt.Println(err)
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user := new(models.User)
	user.FirstName = registerRequest.FirstName
	user.LastName = registerRequest.LastName
	user.Email = registerRequest.Email
	user.Password = registerRequest.Password

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
