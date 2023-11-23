package service

import (
	"errors"

	"github.com/lucasshuan/b2c-ecommerce-api/internal/auth"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/util"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Login(user *model.User, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !util.CheckPasswordHash(password, user.HashedPassword) {
		return "", errors.New("invalid password")
	}

	token, err := auth.GenerateToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) ListAllUsers() ([]*model.User, error) {
	return s.repo.ListAllUsers()
}

func (s *UserService) FindUserByID(userID uint) (*model.User, error) {
	return s.repo.FindUserByID(userID)
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

func (s *UserService) UpdateUser(user *model.User, password string) error {
	if password != "" {
		hashedPassword, err := util.EncryptPassword(password)
		if err != nil {
			return err
		}
		user.HashedPassword = string(hashedPassword)
	}
	return s.repo.SaveUser(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.repo.DeleteUser(user)
}
