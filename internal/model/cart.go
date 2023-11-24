package model

import (
	"encoding/json"
	"io"
)

type Cart struct {
	Base
	IsActive bool
	UserID   uint
	User     User
	Orders   []*Order
	Products []*CartProduct
}

func (c *Cart) UnmarshalGQL(v interface{}) error {
	return nil
}

func (c Cart) MarshalGQL(w io.Writer) {
	c.Base.MarshalGQL(w)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":        c.ID,
		"is_active": c.IsActive,
		"user_id":   c.UserID,
		"user":      c.User,
		"orders":    c.Orders,
		"products":  c.Products,
	})
}
