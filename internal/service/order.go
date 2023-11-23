package service

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(r *repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: r,
	}
}

func (s *OrderService) SaveOrder(order *model.Order) error {
	return s.repo.SaveOrder(order)
}

func (s *OrderService) DeleteOrder(order *model.Order) error {
	return s.repo.DeleteOrder(order)
}

func (s *OrderService) FindOrderByID(id uint) (*model.Order, error) {
	return s.repo.FindOrderByID(id)
}

func (s *OrderService) ListOrdersByUserID(userID uint) ([]*model.Order, error) {
	return s.repo.ListOrdersByUserID(userID)
}
