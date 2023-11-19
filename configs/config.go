package configs

import (
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/spf13/viper"
)

type Tenant struct {
	ID          string
	DatabaseURL string
}

type AppConfig struct {
	Port    string
	Tenants []Tenant
}

func Load() *AppConfig {
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

	var config = &AppConfig{}

	switch env {
	case "dev", "prod":
		config.Port = viper.GetString(env + ".port")
		if config.Port == "" {
			config.Port = "8080"
		}

		tenantsMap := make(map[string]string)
		if err := viper.UnmarshalKey(env+".tenants", &tenantsMap); err != nil {
			log.AppLogger.Fatalf("Failed to unmarshal tenants: %v", err)
		}
		for id, databaseURL := range tenantsMap {
			config.Tenants = append(config.Tenants, Tenant{ID: id, DatabaseURL: databaseURL})
		}
	default:
		log.AppLogger.Fatalf("Unknown environment: %s", env)
	}

	log.AppLogger.Infof("Using %s environment configuration", env)

	return config
}
