package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestLoadEnv(t *testing.T) {
	fmt.Println("Test Case 1 - Key exists")
	key := "port"
	value := LoadEnv(key)
	if value != "8080" {
		t.Errorf("Expected '8080' but found" + value)
	}
	fmt.Println("Test Case 2 - Key does not exist")
	key = "porters"
	value = LoadEnv(key)
	if value != "" {
		t.Errorf("Expected '' but found" + value)
	}
}

func TestExtractApiKey(t *testing.T) {
	fmt.Println("Test Case 1 - Empty Header")
	expected := ""
	actual, err := ExtractAPIKey("")
	if err == nil {
		t.Errorf("Expected error but found " + actual)
	}
	fmt.Println("Test Case 2 - ApiKey exists in header")
	expected = "ApiKey erbaj7ia8nasdasd7"
	actual, err = ExtractAPIKey(expected)
	if err != nil {
		t.Errorf("Expected 'nil' but found " + actual)
	}
}

func TestConvertTimeSQL(t *testing.T) {
	fmt.Println("Test Case 1 - zero time value")
	sqlTime := ConvertTimeSQL(time.Time{})
	if sqlTime.Valid {
		t.Errorf("Expecting 'false' but found 'true'")
	}
	fmt.Println("Test Case 2 - valid time value")
	sqlTime = ConvertTimeSQL(time.Now().UTC())
	if !sqlTime.Valid {
		t.Errorf("Expecting 'true' but found 'false'")
	}
}

func TestConvertStringToTime(t *testing.T) {
	fmt.Println("Test Case 1 - invalid time value")
	sqlTime := ConvertStringToTime("")
	if !sqlTime.IsZero() {
		t.Errorf("Expecting 'true' but found 'false'")
	}
	fmt.Println("Test Case 2 - valid time value")
	sqlTime = ConvertStringToTime("Sun, 23 Jul 2023 10:00:48 +0000")
	if sqlTime.IsZero() {
		t.Errorf("Expecting 'false' but found 'true'")
	}
}
