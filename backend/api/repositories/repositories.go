// Package repositories
package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const queryTimeout = 10 * time.Second

type Repositories struct {
	User IUserRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		User: &userRepository{coll: db.Collection("users")},
	}
}
