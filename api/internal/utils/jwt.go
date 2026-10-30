package utils

import (
	"time"

	"Server/internal/config"
	"Server/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	now := time.Now()

	jwtClaims := JWTClaims{
		UserId: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-api",
			Subject:   user.ID.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(config.Env.JWT.ExpiresAt)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err := token.SignedString(config.Env.JWT.Secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
