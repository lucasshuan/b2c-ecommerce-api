package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	IsActive bool
	Value    float64
	UserID   uint
	User     User
	Orders   []*Order
	Products []*CartProduct
}
