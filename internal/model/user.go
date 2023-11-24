package model

import (
	"encoding/json"
	"io"
)

type User struct {
	Base
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	HashedPassword string `json:"-"`
	IsAdmin        bool   `json:"is_admin"`
	Carts          []*Cart
	Orders         []*Order
}

func (u *User) UnmarshalGQL(v interface{}) error {
	return nil
}

func (u User) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(u)
}
