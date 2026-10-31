package handlers

import (
	"Server/internal/models"
	"Server/internal/services"
	"Server/internal/utils"

	"github.com/gofiber/fiber/v3"
)

var _ IUserHandler = (*userHandler)(nil)

type IUserHandler interface {
	FindOne(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
}

type userHandler struct {
	svc *services.Services
}

// FindOne
// @Summary Find a user
// @Description Retrieve a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Retrieve a user by ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]any
// @Router /api/users/{id} [get]
func (h *userHandler) FindOne(c fiber.Ctx) error {
	userId := c.Params("id")

	user, err := h.svc.User.FindOne(userId)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, user)
}

// Update
// @Summary Update a user
// @Description Update a user information
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User login details"
// @Param request body models.UpdateUser true "User login details"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]any
// @Router /api/users/{id} [patch]
func (h *userHandler) Update(c fiber.Ctx) error {
	userId := c.Params("id")
	updates := new(models.UpdateUser)
	if err := c.Bind().Body(updates); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.svc.User.Update(userId, updates)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, user)
}
