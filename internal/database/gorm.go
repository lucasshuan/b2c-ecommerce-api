package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nobeluc/ecommerce-api/internal/gql"
	"github.com/nobeluc/ecommerce-api/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Databases = make(map[string]*gorm.DB)
	Resolvers = make(map[string]*gql.Resolver)
)

var (
	tenantIDsList = []string{
		"company-01",
		"company-02",
	}
	models = []interface{}{
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Cart{},
		&model.CartProduct{},
		&model.Order{},
	}
)

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	for _, tenantID := range tenantIDsList {
		db, err := initializeGorm(tenantID)
		if err != nil {
			log.Fatalf("failed to initialize database for tenant %s: %v", tenantID, err)
		}
		Databases[tenantID] = db

		resolver := gql.NewResolver(db)
		Resolvers[tenantID] = resolver
	}
}

func initializeGorm(tenantID string) (*gorm.DB, error) {
	dsn := os.Getenv(tenantID + "_MYSQL_DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("dsn not found for tenant: %s", tenantID)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %v", err)
	}

	return db, nil
}
