package main

import (
	"log"
	"os"

	"github.com/gunelmirzoeva/team-tasks/internal/config"
	"github.com/gunelmirzoeva/team-tasks/internal/db"
	"github.com/gunelmirzoeva/team-tasks/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger.Init(cfg.AppEnv)

	logger.Info("Configuration loaded", "port", cfg.Port, "env", cfg.AppEnv)
	logger.Info("Server starting on port " + cfg.Port + "...")

	database, err := db.Connect(cfg.DatabaseURL) 
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer database.Close()
	logger.Info("Database connection established")
}
