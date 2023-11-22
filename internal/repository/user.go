package repository

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) ListAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindUserById(id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SaveUser(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(user *model.User) error {
	return r.DB.Delete(user).Error
}
