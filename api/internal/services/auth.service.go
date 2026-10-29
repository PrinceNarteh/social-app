package services

import (
	"errors"
	"time"

	"Server/internal/models"
	"Server/internal/repositories"
	"Server/internal/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	_                     IAuthService = (*authService)(nil)
	ErrInvalidCredentials              = errors.New("invalid credentials")
)

type IAuthService interface {
	Login(data *models.LoginRequest) (*models.AuthResponse, error)
	Register(user *models.User) (*models.AuthResponse, error)
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Login(data *models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.repo.User.FindByEmail(data.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *authService) Register(user *models.User) (*models.AuthResponse, error) {
	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now

	if err := s.repo.User.Create(user); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, ErrEmailExists
		}
		return nil, err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
