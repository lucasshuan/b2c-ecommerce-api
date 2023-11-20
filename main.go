package main

import (
	"github.com/nobeluc/ecommerce-api/cli/actions"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "b2c-ecommerce-api"
	app.Usage = "Command-line tool for managing the B2C E-commerce API"

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run B2C E-commerce API server",
			Action:  actions.ServerAction,
		},
		{
			Name:    "docker-up",
			Aliases: []string{"u"},
			Usage:   "Run docker containers",
			Action:  actions.DockerUpAction,
		},
		{
			Name:    "docker-down",
			Aliases: []string{"d"},
			Usage:   "Run docker containers",
			Action:  actions.DockerDownAction,
		},
	}

	log.AppLogger.Info("B2C E-commerce API command-line tool started")
}
