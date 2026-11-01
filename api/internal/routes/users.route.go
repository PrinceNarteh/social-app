package routes

import (
	"Server/internal/handlers"
	"Server/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func userRoutes(r fiber.Router, h *handlers.Handlers) {
	routes := r.Group("/users")
	routes.Get("/:id", h.User.FindOne)
	routes.Patch("/:id", middleware.AuthGuard, h.User.Update)
	routes.Delete("/:id", h.User.Delete)
}
