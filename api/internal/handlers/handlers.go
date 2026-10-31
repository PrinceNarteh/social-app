// Package handlers
package handlers

import "Server/internal/services"

type Handlers struct {
	Auth IAuthHandler
	User IUserHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &authHandler{svc: svc},
		User: &userHandler{svc: svc},
	}
}
