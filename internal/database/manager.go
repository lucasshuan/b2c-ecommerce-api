package database

import (
	"net/url"

	"github.com/nobeluc/ecommerce-api/internal/database/scheme"
	"gorm.io/gorm"
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
	dialector := d.getScheme().GetDialector()
	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(d.models...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DBManager) getScheme() scheme.DBScheme {
	switch d.url.Scheme {
	case "mysql":
		return scheme.NewMySQL(d.url)
	case "postgres":
		return scheme.NewPostgres(d.url)
	case "sqlite":
		return scheme.NewSQLite(d.url)
	}
	return nil
}
