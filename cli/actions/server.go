package actions

import (
	"github.com/nobeluc/ecommerce-api/configs"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/server"
	"github.com/urfave/cli"
)

func ServerAction(ctx *cli.Context) error {
	log.AppLogger.Info("Starting server...")
	defer log.AppLogger.Info("Server stopped")

	c := configs.Load()

	database.Init(c)
	server.Start(c)
}
