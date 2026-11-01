package handlers

import (
	"errors"

	"Server/internal/models"
	"Server/internal/services"
	"Server/internal/utils"

	"github.com/gofiber/fiber/v3"
)

var _ IAuthHandler = (*authHandler)(nil)

type IAuthHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}

type authHandler struct {
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
func (h *authHandler) Login(c fiber.Ctx) error {
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
func (h *authHandler) Register(c fiber.Ctx) error {
	registerRequest := new(models.RegisterRequest)
	if err := c.Bind().Body(registerRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "request body not present")
	}

	if err := registerRequest.Validate(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user := &models.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     registerRequest.Email,
		Followers: make([]string, 0),
		Following: make([]string, 0),
	}

	if err := user.HashPassword(); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	authResponse, err := h.svc.Auth.Register(user)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, authResponse)
}
