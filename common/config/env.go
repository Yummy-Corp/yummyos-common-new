package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	envMode := os.Getenv("APP_ENV")
	if envMode == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
