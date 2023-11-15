package server

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/middleware"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	return (func(c *gin.Context) {
		v, ok := c.Get("db")
		if !ok {
			c.AbortWithStatus(500)
			return
		}
		db := v.(*gorm.DB)
		if db == nil {
			c.AbortWithStatus(500)
			return
		}
		resolver := &gql.Resolver{
			DB: db,
		}
		h := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: resolver}))

		h.ServeHTTP(c.Writer, c.Request)
	})
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Start() {
	r := gin.Default()

	r.Use(middleware.TenantMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	log.Printf("Server started on http://localhost:%s", defaultPort)
	log.Printf("Connect to http://localhost:%s/playground for GraphQL playground", defaultPort)
	r.Run(":" + defaultPort)
}
