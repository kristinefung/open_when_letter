package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload" // Do not remove this line, this is to run godotenv.Load() for auto load .env
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

var requiredEnv = []string{
	"DB_USER",
	"DB_PASSWORD",
	"DB_HOST",
	"DB_PORT",
	"DB_NAME",
	"SERVER_PORT",
}

func Load() *Config {
	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			panic(fmt.Sprintf("Required environment variable %s is missing", key))
		}
	}

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
