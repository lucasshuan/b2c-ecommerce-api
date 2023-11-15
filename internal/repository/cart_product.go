package repository

import (
	"multitenant_go_api/internal/model"

	"gorm.io/gorm"
)

type CartProductRepository struct {
	DB *gorm.DB
}

func NewCartProductRepository(db *gorm.DB) *CartProductRepository {
	return &CartProductRepository{
		DB: db,
	}
}

func (r *CartProductRepository) FindByID(id uint) (*model.CartProduct, error) {
	var cartProduct model.CartProduct
	if err := r.DB.First(&cartProduct, id).Error; err != nil {
		return nil, err
	}
	return &cartProduct, nil
}

func (r *CartProductRepository) Save(cartProduct *model.CartProduct) error {
	return r.DB.Save(cartProduct).Error
}

func (r *CartProductRepository) Delete(cartProduct *model.CartProduct) error {
	return r.DB.Delete(cartProduct).Error
}
