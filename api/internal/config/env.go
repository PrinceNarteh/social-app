package config

import (
	"fmt"
	"os"
	"time"
)

var Env = initConfig()

func initConfig() *Config {
	return &Config{
		App: appConfig{
			Port:    fmt.Sprintf(":%s", getEnv("APP_PORT", "4000")),
			Version: getEnv("APP_VERSION", "0.0.1"),
		},
		DB: dbConfig{
			Username: getEnv("DB_USERNAME", "admin"),
			Password: getEnv("DB_PASSWORD", "root"),
			Name:     getEnv("DB_NAME", "test"),
			URL:      getEnv("DB_URL", ""),
		},
		JWT: jwtConfig{
			Secret: []byte(getEnv("JWT_SECRET", "ce26a4a7bead33e504b03f4e13f88015c237f039fe336fec8344bade7d5b0f50e")),
			TTL:    getEnvAsDuration("JWT_TTL", 15*time.Minute),
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
