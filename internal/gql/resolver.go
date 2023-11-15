package gql

import (
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		DB: db,
	}
}
