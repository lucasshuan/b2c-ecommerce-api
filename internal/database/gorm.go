package database

import (
	"fmt"
	"os"

	"github.com/nobeluc/ecommerce-api/internal/logger"
	"github.com/nobeluc/ecommerce-api/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Databases = make(map[string]*gorm.DB)

var (
	tenantIDsList = []string{
		"COMPANY_01",
		"COMPANY_02",
	}
	models = []interface{}{
		&model.Category{},
		&model.User{},
		&model.Product{},
		&model.CartProduct{},
		&model.Order{},
		&model.Cart{},
	}
)

func Initialize() {
	for _, tenantID := range tenantIDsList {
		dsn := os.Getenv(tenantID + "_MYSQL_DATABASE_URL")
		if dsn == "" {
			logger.Log.Fatalf("DSN not found for tenant: %s", tenantID)
		}

		db, err := initializeGorm(dsn)
		if err != nil {
			logger.Log.Fatalf("Failed to initialize database for tenant %s: %v", tenantID, err.Error())
		}

		Databases[tenantID] = db
		println(Databases[tenantID])
	}
}

func initializeGorm(dsn string) (*gorm.DB, error) {
	var datetimePrecision = 2
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn + "?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,                // Add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // Disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // Default datetime precision
		DontSupportRenameIndex:    true,               // Drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // Smart configure based on used version
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %v", err)
	}

	return db, nil
}
