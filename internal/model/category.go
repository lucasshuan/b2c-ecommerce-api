package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string     `gorm:"unique"`
	Products []*Product `gorm:"ForeignKey:CategoryID"`
}
