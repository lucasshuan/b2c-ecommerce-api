package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, loginInput gql.LoginInput) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input gql.CreateUserInput) (*gql.User, error) {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err := r.userService.CreateUser(user, input.Password)
	if err != nil {
		return nil, err
	}

	return &gql.User{
		ID:        strconv.Itoa(int(user.ID)),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input gql.UpdateUserInput) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*gql.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user, err := r.userService.FindUserByID(uint(userID))
	if err != nil {
		return nil, err
	}

	return &gql.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*gql.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}
