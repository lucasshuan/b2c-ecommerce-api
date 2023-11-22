package validation

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/auth"
)

func IsAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	permission := ctx.Value("permission").(auth.Permission)
	if permission != auth.Admin {
		return nil, fmt.Errorf("unauthorized")
	}
	return next(ctx)
}
