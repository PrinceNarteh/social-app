// Package models
package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
	Email     string        `bson:"email" json:"email"`
	Roles     []string      `bson:"roles" json:"roles"`
	Password  string        `bson:"password" json:"-"`
	IsActive  bool          `bson:"isActive" json:"isActive"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

type TokenPair struct {
	AccessToken     string    `json:"accessToken"`
	RefreshToken    string    `json:"refreshToken"`
	RefreshTokenTTL time.Time `json:"-"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 50)),
	)
}

func (u *User) HashPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedBytes)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
