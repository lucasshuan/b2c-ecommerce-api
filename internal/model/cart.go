package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	IsActive bool
	UserID   uint
	User     User
	Orders   []*Order
	Products []*CartProduct
}
