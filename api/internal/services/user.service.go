package services

import (
	"errors"

	"Server/internal/models"
	"Server/internal/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	_                 IUsersService = (*usersService)(nil)
	ErrUserNotFound                 = errors.New("user not found")
	ErrInternalServer               = errors.New("something went wrong")
)

type IUsersService interface {
	FindOne(id string) (*models.User, error)
	Update(id string, updates *models.UpdateUser) (*models.User, error)
	Delete(id string) (string, error)
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
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrUserNotFound
		default:
			return nil, ErrInternalServer
		}
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
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrUserNotFound
		default:
			return nil, ErrInternalServer
		}
	}

	return user, nil
}

func (s *usersService) Delete(id string) (string, error) {
	userId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("invalid id")
	}

	if err := s.repo.User.Delete(userId); err != nil {
		return "", err
	}

	return "user deleted successfully", nil
}
