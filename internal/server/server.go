package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/gql/resolvers"
	"github.com/nobeluc/ecommerce-api/internal/logger"
	"github.com/nobeluc/ecommerce-api/internal/validation"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	return (func(c *gin.Context) {
		// db, ok := c.Get("db")
		// if !ok {
		// 	logger.Log.Error("DB not found in context")
		// 	c.AbortWithStatus(500)
		// 	return
		// }
		db := database.Databases["COMPANY_01"]
		if db == nil {
			logger.Log.Error("DB not found")
			c.AbortWithStatus(500)
			return
		}
		resolver := resolvers.NewResolver(db)
		config := gql.Config{
			Resolvers: resolver,
			Directives: gql.DirectiveRoot{
				Length: validation.LengthDirective,
				Email:  validation.EmailDirective,
			},
		}

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

	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())

	logger.Log.Infof("Connect to http://localhost:%s for GraphQL playground", defaultPort)
	r.Run(":" + defaultPort)
}
