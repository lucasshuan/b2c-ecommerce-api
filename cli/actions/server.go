package actions

import (
	"github.com/lucasshuan/b2c-ecommerce-api/configs"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/database"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/server"
	"github.com/urfave/cli"
)

func ServerAction(ctx *cli.Context) error {
	log.AppLogger.Info("Starting server...")
	defer log.AppLogger.Info("Server stopped")

	configs.Init()
	database.Init()
	server.Init()

	return nil
}
