package main

import (
	"log"

	"github.com/gunelmirzoeva/team-tasks/internal/config"
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

	//TODO: start your server here
}
