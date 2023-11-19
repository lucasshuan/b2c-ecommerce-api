package database

import (
	"net/url"

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
		if tenant.DatabaseURL == "" {
			log.AppLogger.Fatalf("Database url is empty")
		}

		log.AppLogger.Debugf("Connecting to database: %s", tenant.ID)

		parsedURL, err := url.Parse(tenant.DatabaseURL)
		if err != nil {
			log.AppLogger.Fatalf("Failed to parse database url: %v", err)
		}

		dm := NewDBManager(parsedURL, models...)
		db, err := dm.GetDB()
		if err != nil {
			log.AppLogger.Fatalf("Failed to connect to database: %v", err)
		}

		Databases[tenant.ID] = db
	}

	if len(Databases) == 0 {
		log.AppLogger.Fatalf("No database found")
	}
}
