package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/lucasshuan/b2c-ecommerce-api/internal/gql"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input gql.CreateOrderInput) (*gql.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, id string) (*gql.Order, error) {
	panic(fmt.Errorf("not implemented: DeleteOrder - deleteOrder"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*gql.Order, error) {
	panic(fmt.Errorf("not implemented: Order - order"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, userID string) ([]*gql.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// OrderPaymentUpdate is the resolver for the orderPaymentUpdate field.
func (r *subscriptionResolver) OrderPaymentUpdate(ctx context.Context, orderID string) (<-chan *gql.Order, error) {
	panic(fmt.Errorf("not implemented: OrderPaymentUpdate - orderPaymentUpdate"))
}
