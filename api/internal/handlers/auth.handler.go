package handlers

import (
	"Server/internal/models"
	"Server/internal/services"
	"Server/internal/utils"
	"errors"

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

// Login
// @Summary Login a user
// @Description Login a user by providing email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "User login details"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]any
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.Bind().Body(loginRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := loginRequest.Validate(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	authResponse, err := h.svc.Auth.Login(loginRequest)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, authResponse)
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

	if err := registerRequest.Validate(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user := new(models.User)
	user.FirstName = registerRequest.FirstName
	user.LastName = registerRequest.LastName
	user.Email = registerRequest.Email
	user.Followers = make([]string, 0)
	user.Following = make([]string, 0)

	if err := user.HashPassword(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	authResponse, err := h.svc.Auth.Register(user)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, authResponse)
}
