package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/configs"
	"github.com/nobeluc/ecommerce-api/internal/handler"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/middleware"
)

func Init(c *configs.AppConfig) {
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())
	registerRoutes(r)

	if err := r.Run(":" + c.Port); err != nil {
		log.AppLogger.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	r.POST("/:tenantID/graphql", handler.GraphqlHandler())
	r.GET("/:tenantID/playground", handler.PlaygroundHandler())
}
