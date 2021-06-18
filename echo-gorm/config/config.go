package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Driver  string
	Storage string
}

type Config struct {
	DB DbConfig
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env file")
	}

	return &Config{
		DB: DbConfig{
			Driver:  os.Getenv("DB_DRIVER"),
			Storage: os.Getenv("DB_STORAGE"),
		},
	}
}
