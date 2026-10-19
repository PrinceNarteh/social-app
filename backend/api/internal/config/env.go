package config

import (
	"os"
	"time"
)

var Env = initConfig()

func initConfig() *Config {
	return &Config{
		App: appConfig{
			Port:    getEnv("APP_PORT", ":4000"),
			Version: getEnv("APP_VERSION", "0.0.1"),
		},
		DB: dbConfig{
			URI:      getEnv("DB_URI", ""),
			Username: getEnv("DB_USERNAME", "admin"),
			Password: getEnv("DB_PASSWORD", "root"),
			Name:     getEnv("DB_NAME", "test"),
		},
		JWT: jwtConfig{
			AccessTokenSecret:  []byte(getEnv("JWT_ACCESS_TOKEN_SECRET", "ce26a4a7bead33e504b03f4e13f88015c237f039fe336fec8344bade7d5b0f50e")),
			AccessTokenTTL:     getEnvAsDuration("JWT_ACCESS_TOKEN_TTL", 15*time.Minute),
			RefreshTokenSecret: []byte(getEnv("JWT_REFRESH_TOKEN_SECRET", "359039766ear7b8a6e96e6e61feeds1cf46ff217bed35fad5558d919b459f8b4")),
			RefreshTokenTTL:    getEnvAsDuration("JWT_REFRESH_TOKEN_TTL", 7*24*time.Hour),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	timeAsStr := getEnv(key, "")

	if timeAsStr == "" {
		return fallback
	}

	duration, err := time.ParseDuration(timeAsStr)
	if err != nil {
		return fallback
	}

	return duration
}
