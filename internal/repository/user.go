package repository

import (
	"multitenant_go_api/internal/model"

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

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Save(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *model.User) error {
	return r.DB.Delete(user).Error
}
