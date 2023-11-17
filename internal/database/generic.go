package database

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getGormDialector(scheme string, dsn string) gorm.Dialector {
	switch scheme {
	case "mysql":
		return mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "postgres":
		return postgres.New(postgres.Config{
			DSN: dsn,
		})
	case "sqlite", "file":
		return sqlite.Dialector{
			DSN: dsn,
		}
	default:
		return nil
	}
}

func getDSN(url *url.URL) string {
	password, _ := url.User.Password()
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		url.User.Username(),
		password,
		url.Scheme,
		url.Host,
		url.Path[1:], // Remove the leading '/'
	)
	return dsn
}
