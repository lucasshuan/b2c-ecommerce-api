package model

import (
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	Quantity  uint
	CartID    uint
	Cart      Cart
	ProductID uint
	Product   Product
}
