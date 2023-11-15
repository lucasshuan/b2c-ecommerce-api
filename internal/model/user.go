package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	Email          string
	HashedPassword string
	PaymentID      string
	UserCarts      []Cart
}
