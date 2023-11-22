package configs

import (
	"github.com/lucasshuan/b2c-ecommerce-api/internal/log"
	"github.com/spf13/viper"
)

type Tenant struct {
	ID          string
	DatabaseURL string
}

type AppConfig struct {
	JWTSecret string
	Port      string
	Tenants   []Tenant
}

var Config *AppConfig

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.AppLogger.Fatalf("Failed to read configuration file: %v", err)
	}

	env := viper.GetString("env")
	if env == "" {
		env = "dev"
	}

	Config = &AppConfig{}

	switch env {
	case "dev", "prod":
		Config.JWTSecret = viper.GetString(env + ".jwt_secret")
		if Config.JWTSecret == "" {
			log.AppLogger.Fatalf("JWT secret is empty")
		}
		Config.Port = viper.GetString(env + ".port")
		if Config.Port == "" {
			Config.Port = "8080"
		}

		tenantsMap := make(map[string]string)
		if err := viper.UnmarshalKey(env+".tenants", &tenantsMap); err != nil {
			log.AppLogger.Fatalf("Failed to unmarshal tenants: %v", err)
		}
		for id, databaseURL := range tenantsMap {
			Config.Tenants = append(Config.Tenants, Tenant{ID: id, DatabaseURL: databaseURL})
		}
	default:
		log.AppLogger.Fatalf("Unknown environment: %s", env)
	}

	log.AppLogger.Infof("Using %s environment configuration", env)
}
