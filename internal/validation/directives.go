package validation

import (
	"context"
	"fmt"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

func LengthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, min *int, max *int) (res interface{}, err error) {
	fieldValue, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf("invalid field value")
	}

	if min != nil && len(fieldValue) < *min {
		return nil, fmt.Errorf("field value is shorter than the minimum length")
	}

	if max != nil && len(fieldValue) > *max {
		return nil, fmt.Errorf("field value is longer than the maximum length")
	}

	return next(ctx)
}

func EmailDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fieldValue, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf("invalid field value")
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, fieldValue)

	if !match {
		return nil, fmt.Errorf("field value is not a valid email")
	}

	return next(ctx)
}
