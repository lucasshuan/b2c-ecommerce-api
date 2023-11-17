package database

import (
	"fmt"
	"net/url"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dbUrl string, models ...interface{}) (*Database, error) {
	parsedURL, err := url.Parse(dbUrl)
	if err != nil {
		return nil, err
	}

	dsn := getDSN(parsedURL)
	if dsn == "" {
		return nil, fmt.Errorf("dsn is empty")
	}
	dialector := getGormDialector(parsedURL.Scheme, dsn)
	if dialector == nil {
		return nil, fmt.Errorf("dialector %s is not supported", parsedURL.Scheme)
	}

	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
