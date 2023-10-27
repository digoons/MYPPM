package config

import (
	"os"
)

func LoadConfig() *AppConfig {
	return &AppConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		//Load other configuration settings
	}
}
