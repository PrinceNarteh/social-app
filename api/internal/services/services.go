// Package services
package services

import (
	"errors"

	"Server/internal/repositories"
)

var ErrEmailExists = errors.New("email already in used")

type Services struct {
	Auth IAuthService
	User IUsersService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: repo},
		User: &usersService{repo: repo},
	}
}
