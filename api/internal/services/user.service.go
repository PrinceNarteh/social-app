package services

import (
	"errors"

	"Server/internal/models"
	"Server/internal/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ IUsersService = (*usersService)(nil)

type IUsersService interface {
	FindOne(id string) (*models.User, error)
	Update(id string, updates *models.UpdateUser) (*models.User, error)
}

type usersService struct {
	repo *repositories.Repositories
}

func (s *usersService) FindOne(id string) (*models.User, error) {
	userId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	user, err := s.repo.User.FindById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) Update(id string, updates *models.UpdateUser) (*models.User, error) {
	userId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	updateData := bson.M{}

	if updates.FirstName != nil {
		updateData["firstName"] = updates.FirstName
	}

	if updates.LastName != nil {
		updateData["lastName"] = updates.LastName
	}

	if updates.Bio != nil {
		updateData["bio"] = updates.Bio
	}

	if updates.Bio != nil {
		updateData["imageUrl"] = updates.ImageUrl
	}

	user, err := s.repo.User.Update(userId, updateData)
	if err != nil {
		return nil, err
	}

	return user, nil
}
