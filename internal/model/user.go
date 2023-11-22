package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"unique"`
	Email             string
	HashedPassword    string
	OAuthProvider     string
	OAuthAccessToken  string
	OAuthRefreshToken string
	IsAdmin           bool
	Carts             []*Cart
	Orders            []*Order
}
