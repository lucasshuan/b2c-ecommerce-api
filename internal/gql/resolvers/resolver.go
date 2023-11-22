package resolvers

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/service"
	"gorm.io/gorm"
)

type Resolver struct {
	UserService *service.UserService
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		UserService: service.NewUserService(repository.NewUserRepository(db)),
	}
}
