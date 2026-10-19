package repositories

import (
	"context"

	"Server/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ IUserRepository = (*userRepository)(nil)

type IUserRepository interface {
	Find() ([]*models.User, error)
	FindById(id bson.ObjectID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Delete(id bson.ObjectID) error
}

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) findBy(filter any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	user := new(models.User)
	if err := r.coll.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	res, err := r.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(bson.ObjectID)
	return nil
}

func (r *userRepository) FindById(id bson.ObjectID) (*models.User, error) {
	return r.findBy(bson.M{"_id": id})
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	return r.findBy(bson.M{"email": email})
}

func (r *userRepository) Find() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	var users []*models.User
	for cursor.Next(ctx) {
		user := new(models.User)
		if err := cursor.Decode(user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Delete(id bson.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	_, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
