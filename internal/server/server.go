package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasshuan/b2c-ecommerce-api/configs"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/auth"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/handler"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
)

func Init() {
	gin.DefaultWriter = log.GetCustomGinWriter()
	r := gin.Default()

	r.Use(auth.Middleware())
	r.Use(log.Middleware())

	registerRoutes(r)

	if err := r.Run(":" + configs.Config.Port); err != nil {
		log.AppLogger.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	r.POST("/graphql", handler.GraphqlHandler())
	r.GET("/playground/:tenantID", handler.PlaygroundHandler())
}
