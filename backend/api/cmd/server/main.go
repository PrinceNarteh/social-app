package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Server/config"
)

func main() {
	app, db := initApp()

	serverError := make(chan error, 1)
	go func() {
		if err := app.Listen(config.Env.App.Port); err != nil {
			serverError <- err
		}
	}()

	signalError := make(chan os.Signal, 1)
	signal.Notify(signalError, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Printf("System Error: %v", err)
	case err := <-signalError:
		log.Printf("Received Shutdown Signal: %v", err)
	}

	log.Panicln("Server is shutting down...")

	if err := db.Disconnect(); err != nil {
		log.Panicf("Error disconnecting database: %v", err)
	}

	if err := app.Shutdown(); err != nil {
		log.Panicf("Error shutting down server: %v", err)
	}

	log.Panicln("Server Shutdown Successfully")
}
