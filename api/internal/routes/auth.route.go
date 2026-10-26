package routes

import (
	"Server/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func authRoutes(r fiber.Router, h *handlers.Handlers) {
	routes := r.Group("/auth")
	routes.Post("/login", h.Auth.Login)
	routes.Post("/register", h.Auth.Register)
}
