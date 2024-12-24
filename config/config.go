package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	DatabaseURL string
	RedisURL    string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return &Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		RedisURL:    viper.GetString("REDIS_URL"),
	}
}
