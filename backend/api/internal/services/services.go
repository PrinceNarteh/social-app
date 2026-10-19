// Package services
package services

import "Server/repositories"

type Services struct {
	Auth IAuthService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: repo},
	}
}
