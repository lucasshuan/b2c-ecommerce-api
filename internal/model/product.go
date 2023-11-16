package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Price       float64
	Description string
	CategoryID  uint
	Category    Category
}
