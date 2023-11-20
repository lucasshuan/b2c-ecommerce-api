package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/gql/resolvers"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/validation"
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
				Length: validation.LengthDirective,
				Email:  validation.EmailDirective,
			},
		}

		h := handler.NewDefaultServer(gql.NewExecutableSchema(config))

		h.ServeHTTP(c.Writer, c.Request)
	})
}
