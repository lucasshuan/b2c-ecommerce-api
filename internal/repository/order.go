package repository

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) ListOrdersByUserID(userID uint) ([]*model.Order, error) {
	var orders []*model.Order
	if err := r.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) FindOrderByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) SaveOrder(order *model.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepository) DeleteOrder(order *model.Order) error {
	return r.DB.Delete(order).Error
}
