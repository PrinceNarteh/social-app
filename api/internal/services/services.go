// Package services
package services

import (
	"Server/internal/repositories"
	"errors"
)

var ErrEmailExists = errors.New("email already in used.")

type Services struct {
	Auth IAuthService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: repo},
	}
}
