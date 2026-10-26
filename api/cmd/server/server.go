package main

import (
	"log"

	"Server/internal/database"
	"Server/internal/handlers"
	"Server/internal/repositories"
	"Server/internal/routes"
	"Server/internal/services"
	"Server/internal/utils"

	_ "Server/docs"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func initApp() (*fiber.App, *database.MongoDBClient) {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowOriginsFunc: func(origin string) bool {
				return true
			},
		},
	))
	app.Use(logger.New())
	app.Use(recover.New())

	// connecting to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	defer func() {
		if err := db.Disconnect(); err != nil {
			log.Fatalf("Error disconnecting database: %v", err)
		}
	}()

	repo := repositories.NewRepositories(db.Database)
	svc := services.NewServices(repo)
	handlers := handlers.NewHandlers(svc)

	routes.InitRoutes(app, handlers)

	// Serve Swagger Documentation
	app.Get("/swagger/*", swaggo.HandlerDefault)

	return app, db
}
