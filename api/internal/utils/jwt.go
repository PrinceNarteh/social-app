package utils

import (
	"errors"
	"fmt"
	"time"

	"Server/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	now := time.Now()

	jwtClaims := jwt.RegisteredClaims{
		Subject:   userId,
		Issuer:    "social-app-api",
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(config.Env.JWT.ExpiresAt)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err := token.SignedString(config.Env.JWT.Secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method: %v", t.Header["alg"])
		}
		return config.Env.JWT.Secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token or validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
