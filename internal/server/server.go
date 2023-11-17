package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/configs"
	"github.com/nobeluc/ecommerce-api/internal/handler"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/middleware"
)

const defaultPort = "8080"

func Init(c *configs.AppConfig) {
	r := gin.Default()

	r.Use(middleware.LoggingMiddleware())

	r.POST("/graphql", handler.GraphqlHandler())
	r.GET("/", handler.PlaygroundHandler())

	log.AppLogger.Infof("Connect to http://localhost:%s for GraphQL playground", defaultPort)
	r.Run(":" + c.Port)
}
