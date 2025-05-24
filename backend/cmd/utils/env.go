package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("/.env")
	if err != nil {
		log.Fatal("Error loading env files")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
