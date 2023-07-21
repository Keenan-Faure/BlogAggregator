package utils

import (
	"database/sql"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const time_format = time.RFC1123Z

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
func ConvertTimeSQL(timeValue time.Time) sql.NullTime {
	if timeValue.IsZero() {
		return sql.NullTime{
			Time:  timeValue,
			Valid: false,
		}
	}
	return sql.NullTime{
		Time:  timeValue,
		Valid: true,
	}
}

// converts a string to time.Time using time_format
func ConvertStringToTime(timeString string) time.Time {
	timeValue, err := time.Parse(time_format, timeString)
	if err != nil {
		return time.Time{}
	}
	return timeValue
}

// converts a string to a sql.NullString object
func ConvertStringToSQL(description string) sql.NullString {
	if description == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: description,
		Valid:  true,
	}
}

// checks if a string contains certain characters word
func CheckStringWithWord(sentence, word string) bool {
	return strings.Contains(sentence, word)
}

// determines which sorting method to use based on query param
func FindSortParam(sortParam string) string {
	if sortParam != strings.ToLower("asc") {
		return "acs"
	}
	return "desc"
}
