package main

import (
	"github.com/joho/godotenv"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/logger"
	"github.com/nobeluc/ecommerce-api/internal/server"
)

func main() {
	logger.Init()
	logger.Log.Info("starting server...")

	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatalf("failed to load .env file: %v", err)
	}

	database.Initialize()
	server.Start()
}
