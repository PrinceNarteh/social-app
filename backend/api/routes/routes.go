package routes

import (
	"Server/handlers"

	"github.com/gofiber/fiber/v3"
)

func InitRoutes(r fiber.Router, h *handlers.Handlers) {
	api := r.Group("/api")
	authRoutes(api, h)
}
