package database

import "gorm.io/gorm"

type DBScheme interface {
	GetDialector() gorm.Dialector
}
