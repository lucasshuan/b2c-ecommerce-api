package repository

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) ListAllCategories() ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) SaveCategory(category *model.Category) error {
	return r.DB.Save(category).Error
}
func (r *CategoryRepository) DeleteCategory(category *model.Category) error {
	return r.DB.Delete(category).Error
}
