package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/gql/resolvers"
	"github.com/nobeluc/ecommerce-api/internal/logger"
	"github.com/nobeluc/ecommerce-api/internal/middleware"
	"github.com/nobeluc/ecommerce-api/internal/repository"
	"github.com/nobeluc/ecommerce-api/internal/service"
	"github.com/nobeluc/ecommerce-api/internal/validation"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	return (func(c *gin.Context) {
		db, ok := c.Get("db")
		if !ok {
			logger.Log.Error("DB not found in context")
			c.AbortWithStatus(500)
			return
		}
		resolver := resolvers.Resolver{
			UserService: service.NewUserService(repository.NewUserRepository(db.(*gorm.DB))),
		}
		config := gql.Config{Resolvers: &resolver}

		config.Directives.Length = validation.LengthDirective
		config.Directives.Email = validation.EmailDirective

		h := handler.NewDefaultServer(gql.NewExecutableSchema(config))

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

	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.TenantMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	logger.Log.Infof("Connect to http://localhost:%s for GraphQL playground", defaultPort)
	r.Run(":" + defaultPort)
}
