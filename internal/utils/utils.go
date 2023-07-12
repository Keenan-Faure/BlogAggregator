package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	godotenv.Load()
	value := os.Getenv(strings.ToUpper(key))
	return value
}
