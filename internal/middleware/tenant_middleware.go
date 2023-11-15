package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"gorm.io/gorm"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("X-Tenant-ID")
		db := getDatabaseForTenant(tenantID)
		if db == nil {
			c.AbortWithStatus(500)
			return
		}

		c.Set("db", db)
		c.Next()
	}
}

func getDatabaseForTenant(tenantID string) *gorm.DB {
	return database.Databases[tenantID]
}
