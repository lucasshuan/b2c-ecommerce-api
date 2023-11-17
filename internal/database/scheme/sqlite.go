package scheme

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLite struct {
	url *url.URL
}

func NewSQLite(url *url.URL) *SQLite {
	return &SQLite{
		url: url,
	}
}

func (s *SQLite) GetDialector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: s.getDSN(),
	})
}

func (s *SQLite) getDSN() string {
	password, _ := s.url.User.Password()
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.url.User.Username(),
		password,
		s.url.Scheme,
		s.url.Host,
		s.url.Path[1:], // Remove the leading '/'
	)
	return dsn
}
