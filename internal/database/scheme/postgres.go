package scheme

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Postgres struct {
	url *url.URL
}

func NewPostgres(url *url.URL) *Postgres {
	return &Postgres{
		url: url,
	}
}

func (s *Postgres) GetDialector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: s.getDSN(),
	})
}

func (s *Postgres) getDSN() string {
	password, _ := s.url.User.Password()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		s.url.User.Username(),
		password,
		s.url.Hostname(),
		s.url.Port(),
		s.url.Path[1:], // Remove the leading '/'
	)
	return dsn
}
