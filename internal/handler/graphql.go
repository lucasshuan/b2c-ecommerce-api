package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/database"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/gql"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/gql/resolvers"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/gql/validation"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
)

func GraphqlHandler() gin.HandlerFunc {
	return (func(c *gin.Context) {
		tenantID := c.GetHeader("X-Tenant-ID")
		if tenantID == "" {
			log.AppLogger.Error("TenantID not found in context")
			c.Abort()
			return
		}
		db := database.Databases[tenantID]
		if db == nil {
			log.AppLogger.Error("Database not found for tenantID: %s", tenantID)
			c.Abort()
			return
		}
		resolver := resolvers.NewResolver(db)
		config := gql.Config{
			Resolvers: resolver,
			Directives: gql.DirectiveRoot{
				Length:  validation.Length,
				Email:   validation.Email,
				IsUser:  validation.IsUser,
				IsAdmin: validation.IsAdmin,
			},
		}

		h := handler.NewDefaultServer(gql.NewExecutableSchema(config))

		h.ServeHTTP(c.Writer, c.Request)
	})
}
