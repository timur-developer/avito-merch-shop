package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
}

func Load() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/merch_shop?sslmode=disable")
	viper.SetDefault("JWT_SECRET", "secret-jwt-key-change-in-prod")

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return cfg
}
