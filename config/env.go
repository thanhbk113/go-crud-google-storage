package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvGet(name string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	return os.Getenv(name)
}
