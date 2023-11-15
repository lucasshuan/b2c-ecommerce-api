package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Value          float64
	AsaasPaymentID string
	CartID         uint
	UserID         uint
	Cart           Cart
	User           User
}
