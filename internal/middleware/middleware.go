package middleware

import (
	"context"
	"net/http"

	"github.com/nobeluc/ecommerce-api/internal/database"
	"gorm.io/gorm"
)

func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")
		db := getDatabaseForTenant(tenantID)
		// Defina o banco de dados correto no contexto da solicitação
		ctx := context.WithValue(r.Context(), "db", db)
		// Chame o próximo manipulador com o contexto atualizado
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getDatabaseForTenant(tenantID string) *gorm.DB {
	return database.Databases[tenantID]
}
