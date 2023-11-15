package repository

import (
	"multitenant_go_api/internal/model"

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

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Save(category *model.Category) error {
	return r.DB.Save(category).Error
}
func (r *CategoryRepository) Delete(category *model.Category) error {
	return r.DB.Delete(category).Error
}
