package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/lucasshuan/b2c-ecommerce-api/internal/gql"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input gql.CreateCategoryInput) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input gql.UpdateCategoryInput) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id uint) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id uint) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}
