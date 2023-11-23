package service

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(r *repository.CartRepository) *CartService {
	return &CartService{
		repo: r,
	}
}

func (s *CartService) ListCartsByUserID(userID uint) ([]*model.Cart, error) {
	return s.repo.ListCartsByUserID(userID)
}

func (s *CartService) FindCartByID(cartID uint) (*model.Cart, error) {
	return s.repo.FindCartByID(cartID)
}

func (s *CartService) AddProductToCart(cartProduct *model.CartProduct) error {
	return s.repo.SaveCartProduct(cartProduct)
}

func (s *CartService) RemoveProductFromCart(cartProduct *model.CartProduct) error {
	return s.repo.DeleteCartProduct(cartProduct)
}

func (s *CartService) EmptyCart(cart *model.Cart) error {
	cartProduct := &model.CartProduct{
		CartID: cart.ID,
	}
	cart.Products = []*model.CartProduct{}
	return s.repo.DeleteCartProduct(cartProduct)
}
