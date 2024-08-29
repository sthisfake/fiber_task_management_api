package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
