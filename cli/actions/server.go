package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasshuan/b2c-ecommerce-api/configs"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/auth"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/database"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/handler"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
	"github.com/urfave/cli"
)

func ServerAction(ctx *cli.Context) error {
	log.AppLogger.Info("Starting server...")
	defer log.AppLogger.Info("Server stopped")

	configs.Init()
	database.Init()

	gin.DefaultWriter = log.GetCustomGinWriter()
	r := gin.Default()

	r.Use(auth.Middleware())
	r.Use(log.Middleware())

	registerRoutes(r)

	if err := r.Run(":" + configs.Config.Port); err != nil {
		log.AppLogger.Fatalf("Failed to start server: %v", err)
	}

	return nil
}

func registerRoutes(r *gin.Engine) {
	r.POST("/graphql", handler.GraphqlHandler())
	r.GET("/playground/:tenantID", handler.PlaygroundHandler())
}
