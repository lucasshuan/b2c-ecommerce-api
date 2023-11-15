package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	IsActive bool
	Value    float64
	UserID   uint
	OrderID  uint
	User     User
	Order    *Order
	Products []Product
}
