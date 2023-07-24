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
	fmt.Println("Test Case 3 - valid time value - incorrect format")
	sqlTime = ConvertStringToTime("23 Jul 2023 10:00:48")
	if !sqlTime.IsZero() {
		t.Errorf("Expecting 'false' but found 'true'")
	}
}

func TestConvertStringToSQL(t *testing.T) {
	fmt.Println("Test Case 1 - valid string value")
	sqlString := ConvertStringToSQL("easy")
	if !sqlString.Valid {
		t.Errorf("Expecting 'true' but found 'false'")
	}
	fmt.Println("Test Case 2 - invalid time value")
	sqlString = ConvertStringToSQL("")
	if sqlString.Valid {
		t.Errorf("Expecting 'false' but found 'true'")
	}
}

func TestCheckStringWithWord(t *testing.T) {
	fmt.Println("Test Case 1 - string exists in sentence")
	exists := CheckStringWithWord("easy", "e")
	if !exists {
		t.Errorf("Expecting 'true' but found 'false'")
	}
	fmt.Println("Test Case 2 - string DNE in sentence")
	exists = CheckStringWithWord("piece", "one")
	if exists {
		t.Errorf("Expecting 'false' but found 'true'")
	}
}

func TestFindSortParam(t *testing.T) {
	fmt.Println("Test Case 1 - string exists in sentence")
	param := FindSortParam("asc")
	if param != "asc" {
		t.Errorf("Expecting 'asc' but found 'desc'")
	}
	fmt.Println("Test Case 2 - string DNE in sentence")
	param = FindSortParam("")
	if param != "desc" {
		t.Errorf("Expecting 'desc' but found 'asc'")
	}
	fmt.Println("Test Case 3 - string exists in sentence - UPPER")
	param = FindSortParam("ASC")
	if param != "asc" {
		t.Errorf("Expecting 'asc' but found 'desc'")
	}
}

func TestAddSearchChar(t *testing.T) {
	fmt.Println("Test Case 1 - Valid string")
	srt := "pokemon_test"
	expected := "%(" + srt + ")%"
	if AddSearchChar(srt) != expected {
		t.Errorf("Expected " + expected + " but found " + AddSearchChar(srt))
	}
}
