package validation

import (
	"context"
	"fmt"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

func Email(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fieldName := graphql.GetFieldContext(ctx).Field.Name
	fieldValue, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf("invalid %s value", fieldName)
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, fieldValue)

	if !match {
		return nil, fmt.Errorf("%s value is not a valid email", fieldName)
	}

	return next(ctx)
}
