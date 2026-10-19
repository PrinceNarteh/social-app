package main

import (
	"log"

	"Server/database"
	"Server/handlers"
	"Server/repositories"
	"Server/routes"
	"Server/services"
	"Server/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func initApp() (*fiber.App, *database.MongoDBClient) {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())
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

	return app, db
}
