// package gql

// import (
// 	"context"
// 	"strconv"

// 	"github.com/nobeluc/ecommerce-api/internal/model"
// 	"github.com/nobeluc/ecommerce-api/internal/repository"
// 	"github.com/nobeluc/ecommerce-api/internal/service"
// )

// func (r *queryResolver) GetUser(ctx context.Context, id string) (*User, error) {
// }

// func (r *Resolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
// 	hashedPassword, err := service.EncryptPassword(input.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	user := &model.User{
// 		Name:           input.Name,
// 		Email:          input.Email,
// 		HashedPassword: string(hashedPassword),
// 	}
// 	repo := repository.NewUserRepository(r.DB)
// 	err = repo.SaveUser(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res := &User{
// 		ID:    strconv.Itoa(int(user.ID)),
// 		Name:  user.Name,
// 		Email: user.Email,
// 	}
// 	return res, nil
// }
