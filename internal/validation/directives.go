package validation

import (
	"context"
	"fmt"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

func LengthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, min *int, max *int) (res interface{}, err error) {
	fieldName := graphql.GetFieldContext(ctx).Field.Name
	fieldValue, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf("invalid %s value", fieldName)
	}

	if min != nil && len(fieldValue) < *min {
		return nil, fmt.Errorf("%s value is shorter than the minimum length", fieldName)
	}

	if max != nil && len(fieldValue) > *max {
		return nil, fmt.Errorf("%s value is longer than the maximum length", fieldName)
	}

	return next(ctx)
}

func EmailDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
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
