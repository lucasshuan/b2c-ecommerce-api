package configs

import (
	"strings"

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
	env := viper.GetString("env")
	if env == "" {
		env = "dev"
	}

	viper.SetEnvPrefix(env)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.AppLogger.Fatalf("Failed to read configuration file: %v", err)
	}

	var config = AppConfig{}

	switch env {
	case "dev", "prod":
		config.Port = viper.GetString("port")
		viper.UnmarshalKey("tenants", &config.Tenants)
	default:
		log.AppLogger.Fatalf("Unknown environment: %s", env)
	}

	log.AppLogger.Infof("Using %s environment configuration\n", env)

	return &config
}
