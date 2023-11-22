package database

import (
	"net/url"

	"github.com/lucasshuan/b2c-ecommerce-api/internal/database/scheme"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBManager struct {
	url    *url.URL
	models []interface{}
}

func NewDBManager(url *url.URL, models ...interface{}) *DBManager {
	return &DBManager{
		url:    url,
		models: models,
	}
}

func (d *DBManager) GetDB() (*gorm.DB, error) {
	dialector := d.getDBScheme().GetDialector()
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(d.models...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DBManager) getDBScheme() scheme.DBScheme {
	switch d.url.Scheme {
	case "mysql":
		return scheme.NewMySQL(d.url)
	case "postgres":
		return scheme.NewPostgres(d.url)
	default:
		return nil
	}
}
