package model

import (
	"encoding/json"
	"io"
)

type Category struct {
	Base
	Name     string `gorm:"unique"`
	Products []*Product
}

func (c *Category) UnmarshalGQL(v interface{}) error {
	return nil
}

func (c Category) MarshalGQL(w io.Writer) {
	c.Base.MarshalGQL(w)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"name":     c.Name,
		"products": c.Products,
	})
}
