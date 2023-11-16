package main

import (
	"github.com/joho/godotenv"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/logger"
	"github.com/nobeluc/ecommerce-api/internal/server"
)

func main() {
	logger.Init()
	logger.Log.Info("Starting server...")

	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatalf("Failed to load .env file: %v", err)
	}

	database.Initialize()
	server.Start()
}
