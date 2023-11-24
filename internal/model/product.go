package model

type Product struct {
	Base
	Name        string `gorm:"unique"`
	Price       float64
	Description string
	CategoryID  uint
	Category    Category
}
