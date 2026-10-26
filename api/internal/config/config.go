// Package config
package config

import "time"

type Config struct {
	App appConfig
	DB  dbConfig
	JWT jwtConfig
}

type appConfig struct {
	Port    string
	Version string
}

type dbConfig struct {
	URI      string
	Username string
	Password string
	Name     string
}

type jwtConfig struct {
	AccessTokenSecret  []byte
	AccessTokenTTL     time.Duration
	RefreshTokenSecret []byte
	RefreshTokenTTL    time.Duration
}
