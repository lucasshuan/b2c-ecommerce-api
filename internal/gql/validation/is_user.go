package validation

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/auth"
)

func IsUser(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	permission := ctx.Value("permission").(auth.Permission)
	if permission == auth.Guest {
		return nil, fmt.Errorf("unauthorized")
	}
	return next(ctx)
}
