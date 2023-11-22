package main

import (
	"os"

	"github.com/lucasshuan/b2c-ecommerce-api/cli/actions"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "b2c-ecommerce-cli"
	app.Usage = "Command-line tool for managing the B2C E-commerce API"

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start server",
			Action:  actions.ServerAction,
		},
		{
			Name:    "docker-up",
			Aliases: []string{"u"},
			Usage:   "Start docker containers",
			Action:  actions.DockerUpAction,
		},
		{
			Name:    "docker-down",
			Aliases: []string{"d"},
			Usage:   "Stop docker containers",
			Action:  actions.DockerDownAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.AppLogger.Fatalf("Error starting the application: %v", err)
	}
}
