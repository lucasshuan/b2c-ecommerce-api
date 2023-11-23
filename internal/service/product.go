package service

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(r *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: r,
	}
}

func (s *ProductService) ListProducts() ([]*model.Product, error) {
	return s.repo.ListProductsByFilters(map[string]interface{}{})
}

func (s *ProductService) FindProductByID(id uint) (*model.Product, error) {
	return s.repo.FindProductByID(id)
}

func (s *ProductService) SaveProduct(product *model.Product) error {
	return s.repo.SaveProduct(product)
}

func (s *ProductService) DeleteProduct(product *model.Product) error {
	return s.repo.DeleteProduct(product)
}
