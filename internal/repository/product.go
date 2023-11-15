package repository

import (
	"github.com/nobeluc/ecommerce-api/internal/model"
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

func (r *ProductRepository) ListProductsByFilters(filters map[string]interface{}) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Where(filters).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindProductById(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) SaveProduct(product *model.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(product *model.Product) error {
	return r.DB.Delete(product).Error
}
