package handler

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/log"
)

func PlaygroundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantID")
		if tenantID == "" {
			log.AppLogger.Error("TenantID not found in context")
			c.Abort()
			return
		}
		fh := map[string]string{
			"X-Tenant-ID": tenantID,
		}
		h := playground.HandlerWithHeaders("GraphQL playground", "/graphql", fh, nil)
		h.ServeHTTP(c.Writer, c.Request)
	}
}
