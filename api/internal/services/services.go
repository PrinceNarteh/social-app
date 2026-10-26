// Package services
package services

import "Server/internal/repositories"

type Services struct {
	Auth IAuthService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: repo},
	}
}
