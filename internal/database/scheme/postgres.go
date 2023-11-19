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
		DSN:                  s.getDSN(),
		PreferSimpleProtocol: true,
	})
}

func (s *Postgres) getDSN() string {
	password, _ := s.url.User.Password()

	// host=%[1]s user=%[2]s password=%[3]s dbname=%[4]s port=%[5]s sslmode=disable TimeZone=UTC
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		s.url.Hostname(),
		s.url.User.Username(),
		password,
		s.url.Path[1:],
		s.url.Port(),
	)
}
