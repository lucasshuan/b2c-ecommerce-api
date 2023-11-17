package scheme

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
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
	return postgres.New(postgres.Config{
		DSN: s.getDSN(),
	})
}

func (s *Postgres) getDSN() string {
	password, _ := s.url.User.Password()

	// user=%[1]s password=%[2]s dbname=%[3]s sslmode=require host=%[4]s port=%[5]s
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=require host=%s port=%s",
		s.url.User.Username(),
		password,
		s.url.Path[1:],
		s.url.Hostname(),
		s.url.Port(),
	)
	return dsn
}
