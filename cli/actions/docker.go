package actions

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
	"github.com/urfave/cli"
)

func DockerUpAction(ctx *cli.Context) error {
	log.AppLogger.Info("Starting docker containers...")

	composeFiles, err := filepath.Glob("./docker/*/compose.yml")
	if err != nil {
		return err
	}

	for _, composeFile := range composeFiles {
		cmd := exec.Command("docker-compose", "-f", composeFile, "up", "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		cmd.Run()
	}
	return nil
}

func DockerDownAction(ctx *cli.Context) error {
	log.AppLogger.Info("Stopping docker containers...")

	composeFiles, err := filepath.Glob("./docker/*/compose.yml")
	if err != nil {
		return err
	}

	for _, composeFile := range composeFiles {
		cmd := exec.Command("docker-compose", "-f", composeFile, "down", "-v")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	return nil
}
