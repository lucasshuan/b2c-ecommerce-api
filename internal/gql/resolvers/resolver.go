package resolvers

import (
	"github.com/nobeluc/ecommerce-api/internal/repository"
	"github.com/nobeluc/ecommerce-api/internal/service"
	"gorm.io/gorm"
)

type Resolver struct {
	userService *service.UserService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		userService: service.NewUserService(repository.NewUserRepository(db)),
	}
}
