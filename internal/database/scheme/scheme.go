package scheme

import "gorm.io/gorm"

type DBScheme interface {
	GetDialector() gorm.Dialector
}
