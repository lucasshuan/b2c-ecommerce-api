package model

type Order struct {
	Base
	Value          float64
	AsaasPaymentID string
	CartID         uint
	Cart           Cart
	UserID         uint
	User           User
}
