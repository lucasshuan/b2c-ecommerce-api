package repository

import (
	"multitenant_go_api/internal/model"

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

func (r *OrderRepository) FindByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Save(order *model.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepository) Delete(order *model.Order) error {
	return r.DB.Delete(order).Error
}
