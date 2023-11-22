package repository

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
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

func (r *CartProductRepository) SaveCartProduct(cartProduct *model.CartProduct) error {
	return r.DB.Save(cartProduct).Error
}

func (r *CartProductRepository) DeleteCartProduct(cartProduct *model.CartProduct) error {
	return r.DB.Delete(cartProduct).Error
}
