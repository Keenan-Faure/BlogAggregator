package utils

import (
	"database/sql"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	godotenv.Load()
	value := os.Getenv(strings.ToUpper(key))
	return value
}

// Authorization: ApiKey <key>
func ExtractAPIKey(authString string) (string, error) {
	if authString == "" {
		return "", errors.New("no Authorization found in header")
	}
	if len(authString) <= 7 {
		return "", errors.New("malformed Auth Header")
	}
	if authString[0:6] != "ApiKey" {
		return "", errors.New("malformed second part of authentication")
	}
	return authString[7:], nil
}

// convert time.Time to sql.NullTime
func ConvertTimeSQL(time time.Time) sql.NullTime {
	if time.IsZero() {
		return sql.NullTime{
			Time:  time,
			Valid: false,
		}
	}
	return sql.NullTime{
		Time:  time,
		Valid: true,
	}
}
