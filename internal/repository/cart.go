package repository

import (
	"multitenant_go_api/internal/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (r *CartRepository) FindByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	if err := r.DB.First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) Save(cart *model.Cart) error {
	return r.DB.Save(cart).Error
}

func (r *CartRepository) Delete(cart *model.Cart) error {
	return r.DB.Delete(cart).Error
}
