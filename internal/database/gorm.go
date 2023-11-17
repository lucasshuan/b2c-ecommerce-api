package database

import (
	"github.com/nobeluc/ecommerce-api/configs"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/model"
	"gorm.io/gorm"
)

var Databases = make(map[string]*gorm.DB)

func Init(c *configs.AppConfig) {
	models := []interface{}{
		&model.Category{},
		&model.User{},
		&model.Product{},
		&model.CartProduct{},
		&model.Order{},
		&model.Cart{},
	}
	for _, tenant := range c.Tenants {
		db, err := NewDatabase(tenant.DatabaseURL, models...)
		if err != nil {
			log.AppLogger.Fatal(err)
		}
		Databases[tenant.ID] = db.DB
	}
}
