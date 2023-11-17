package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/gql/resolvers"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/validation"
	"gorm.io/gorm"
)

func getDb(c *gin.Context) *gorm.DB {
	tenantId := c.GetHeader("X-Tenant-Id")
	if tenantId == "" {
		return nil
	}
	db := database.Databases[tenantId]
	return db
}

func GraphqlHandler() gin.HandlerFunc {
	return (func(c *gin.Context) {
		db := getDb(c)
		if db == nil {
			log.AppLogger.Error("Database not found")
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
