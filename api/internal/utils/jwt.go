package utils

import (
	"time"

	"Server/internal/config"
	"Server/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	UserId string
	Email  string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func GenerateTokens(user *models.User) (*models.TokenPair, error) {
	now := time.Now()

	accessClaims := AccessClaims{
		UserId: user.ID.Hex(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-api",
			Subject:   user.ID.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(config.Env.JWT.AccessTokenTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := token.SignedString(config.Env.JWT.AccessTokenSecret)
	if err != nil {
		return nil, err
	}

	refreshTokenTTL := now.Add(config.Env.JWT.RefreshTokenTTL)
	refreshClaims := RefreshClaims{
		UserId: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-api",
			Subject:   user.ID.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(refreshTokenTTL),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := token.SignedString(config.Env.JWT.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		RefreshTokenTTL: refreshTokenTTL,
	}, nil
}
