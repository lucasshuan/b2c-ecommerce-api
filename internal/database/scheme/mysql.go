package scheme

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	url *url.URL
}

func NewMySQL(url *url.URL) *MySQL {
	return &MySQL{
		url: url,
	}
}

func (s *MySQL) GetDialector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: s.getDSN(),
	})
}

func (s *MySQL) getDSN() string {
	password, _ := s.url.User.Password()
	// user:password@tcp(host:port)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.url.User.Username(),
		password,
		s.url.Hostname(),
		s.url.Port(),
		s.url.Path[1:],
	)
}
