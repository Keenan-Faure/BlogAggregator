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

// extracts the API Key from the auth string
func ExtractAPIKey(authString string) string {
	if authString == "" {
		return ""
	}
	return authString[7:]
}
