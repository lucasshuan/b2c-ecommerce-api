package repository

import (
	"multitenant_go_api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Save(product *model.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(product *model.Product) error {
	return r.DB.Delete(product).Error
}
