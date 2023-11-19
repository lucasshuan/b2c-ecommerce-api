package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/nobeluc/ecommerce-api/internal/gql"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input gql.CreateCategoryInput) (*gql.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input gql.UpdateCategoryInput) (*gql.Category, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*gql.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*gql.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}
