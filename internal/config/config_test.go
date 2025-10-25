package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/dbname")
	os.Setenv("JWT_SECRET", "mysecretkey")
	os.Setenv("APP_ENV", "development")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}

	if cfg.Port != "8080" {
        t.Errorf("Expected Port to be '8080', got '%s'", cfg.Port)
    }
    if cfg.DatabaseURL != "postgres://user:pass@localhost:5432/dbname" {
        t.Errorf("Expected DatabaseURL to be correct, got '%s'", cfg.DatabaseURL)
    }
    if cfg.JWTSecret != "mysecretkey" {
        t.Errorf("Expected JWTSecret to be 'mysecretkey', got '%s'", cfg.JWTSecret)
    }
    if cfg.AppEnv != "development" {
        t.Errorf("Expected AppEnv to be 'development', got '%s'", cfg.AppEnv)
    }
}