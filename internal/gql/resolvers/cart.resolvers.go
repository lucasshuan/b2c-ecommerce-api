package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/nobeluc/ecommerce-api/internal/gql"
)

// CreateCart is the resolver for the createCart field.
func (r *mutationResolver) CreateCart(ctx context.Context, input gql.CreateCartInput) (*gql.Cart, error) {
	panic(fmt.Errorf("not implemented: CreateCart - createCart"))
}

// AddProductToCart is the resolver for the addProductToCart field.
func (r *mutationResolver) AddProductToCart(ctx context.Context, input gql.AddProductToCartInput) (*gql.CartProduct, error) {
	panic(fmt.Errorf("not implemented: AddProductToCart - addProductToCart"))
}

// RemoveProductFromCart is the resolver for the removeProductFromCart field.
func (r *mutationResolver) RemoveProductFromCart(ctx context.Context, cartProductID string) (*gql.Cart, error) {
	panic(fmt.Errorf("not implemented: RemoveProductFromCart - removeProductFromCart"))
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context, id string) (*gql.Cart, error) {
	panic(fmt.Errorf("not implemented: Cart - cart"))
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context, userID string) ([]*gql.Cart, error) {
	panic(fmt.Errorf("not implemented: Carts - carts"))
}
