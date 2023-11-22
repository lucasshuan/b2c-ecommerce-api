package repository

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
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

func (r *CartRepository) ListCartsByUserID(userID uint) ([]*model.Cart, error) {
	var carts []*model.Cart
	if err := r.DB.Where("user_id = ?", userID).Preload("CartProducts").Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) FindCartByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	if err := r.DB.First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) SaveCart(cart *model.Cart) error {
	return r.DB.Save(cart).Error
}

func (r *CartRepository) DeleteCart(cart *model.Cart) error {
	return r.DB.Delete(cart).Error
}
