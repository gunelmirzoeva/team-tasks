package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	AppEnv      string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		AppEnv:      os.Getenv("APP_ENV"),
	}

	if cfg.Port == "" || cfg.DatabaseURL == "" || cfg.JWTSecret == "" || cfg.AppEnv == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
