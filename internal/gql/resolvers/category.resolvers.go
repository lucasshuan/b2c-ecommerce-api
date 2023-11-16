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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetCategory(ctx context.Context, id string) (*gql.Category, error) {
	panic(fmt.Errorf("not implemented: GetCategory - getCategory"))
}
func (r *queryResolver) ListCategories(ctx context.Context) ([]*gql.Category, error) {
	panic(fmt.Errorf("not implemented: ListCategories - listCategories"))
}
