package service

import (
	"errors"

	"github.com/nobeluc/ecommerce-api/internal/model"
	"github.com/nobeluc/ecommerce-api/internal/repository"
	"github.com/nobeluc/ecommerce-api/internal/util"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) ListAllUsers() ([]*model.User, error) {
	return s.repo.ListAllUsers()
}

func (s *UserService) FindUserById(id uint) (*model.User, error) {
	return s.repo.FindUserById(id)
}

func (s *UserService) CreateUser(user *model.User, password string) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("invalid user")
	}
	hashedPassword, err := util.EncryptPassword(password)
	if err != nil {
		return err
	}
	user.HashedPassword = string(hashedPassword)

	return s.repo.SaveUser(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.SaveUser(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.repo.DeleteUser(user)
}
