package model

import (
	"encoding/json"
	"io"
)

type CartProduct struct {
	Base
	Quantity  uint `json:"quantity"`
	CartID    uint `json:"cart_id"`
	Cart      Cart
	ProductID uint `json:"product_id"`
	Product   Product
}

func (cp *CartProduct) UnmarshalGQL(v interface{}) error {
	return nil
}

func (cp CartProduct) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"quantity":  cp.Quantity,
		"cartID":    cp.CartID,
		"productID": cp.ProductID,
	})
}
