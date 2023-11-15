package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"unique"`
	Email          string
	HashedPassword string
	PaymentID      string
	Carts          []Cart
	Orders         []Order
}
