package services

import (
	"errors"
	"time"

	"Server/internal/models"
	"Server/internal/repositories"
	"Server/internal/utils"
)

var (
	_                     IAuthService = (*authService)(nil)
	ErrInvalidCredentials              = errors.New("invalid credentials")
)

type IAuthService interface {
	Login(data *models.LoginRequest) (*models.User, *models.TokenPair, error)
	Register(user *models.User) (*models.TokenPair, error)
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Login(data *models.LoginRequest) (*models.User, *models.TokenPair, error) {
	user, err := s.repo.User.FindByEmail(data.Email)
	if err != nil {
		return nil, nil, ErrInvalidCredentials
	}

	if !user.VerifyPassword(data.Password) {
		return nil, nil, ErrInvalidCredentials
	}

	tokenPair, err := utils.GenerateTokens(user)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

func (s *authService) Register(user *models.User) (*models.TokenPair, error) {
	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now

	if err := s.repo.User.Create(user); err != nil {
		return nil, err
	}

	tokenPair, err := utils.GenerateTokens(user)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}
