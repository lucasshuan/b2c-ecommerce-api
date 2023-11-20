package actions

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/urfave/cli"
)

func DockerUpAction(ctx *cli.Context) error {
	log.AppLogger.Info("Starting docker containers...")
	defer log.AppLogger.Info("Docker containers created")

	composeFiles, err := filepath.Glob("./docker/*/compose.yml")
	if err != nil {
		return err
	}

	for _, composeFile := range composeFiles {
		cmd := exec.Command("docker-compose", "-f", composeFile, "up", "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			log.AppLogger.Errorf("Error running command for compose file %s: %v", composeFile, err)
			continue
		}
		log.AppLogger.Infof("Successfully started containers from compose file %s", composeFile)
	}
	return nil
}

func DockerDownAction(ctx *cli.Context) error {
	log.AppLogger.Info("Stopping docker containers...")
	defer log.AppLogger.Info("Docker containers stopped")

	composeFiles, err := filepath.Glob("./docker/*/compose.yml")
	if err != nil {
		return err
	}

	for _, composeFile := range composeFiles {
		cmd := exec.Command("docker-compose", "-f", composeFile, "down", "-v")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			log.AppLogger.Errorf("Error running command for compose file %s: %v", composeFile, err)
			continue
		}
		log.AppLogger.Infof("Successfully stopped containers from compose file %s", composeFile)
	}
	return nil
}
