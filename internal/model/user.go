package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"unique"`
	HashedPassword string
	IsAdmin        bool
	Carts          []*Cart
	Orders         []*Order
}
