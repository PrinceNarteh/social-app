// Package handlers
package handlers

import "Server/services"

type Handlers struct {
	Auth IAuthHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &AuthHandler{svc: svc},
	}
}
