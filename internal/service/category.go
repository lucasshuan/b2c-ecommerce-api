package service

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"github.com/lucasshuan/b2c-ecommerce-api/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(r *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: r,
	}
}

func (s *CategoryService) ListCategories() ([]*model.Category, error) {
	return s.repo.ListAllCategories()
}

func (s *CategoryService) SaveCategory(category *model.Category) error {
	return s.repo.SaveCategory(category)
}

func (s *CategoryService) DeleteCategory(category *model.Category) error {
	return s.repo.DeleteCategory(category)
}
