package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI(key string, defaultVal string) string {
	err := godotenv.Load()
	value := os.Getenv(key)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return value
}
