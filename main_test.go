package main

import (
	"blog/internal/database"
	"context"
	"fmt"
	"utils"

	"testing"

	_ "github.com/lib/pq"
)

// Tests below assume that the tables exist
// and the correct database was chosen
// all data created by tests are removed upon finish

func TestDatabaseConnection(t *testing.T) {
	fmt.Println("Test Case 1 - Invalid database url string")
	dbconfig, err := InitConn("abc123")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	_, err = dbconfig.DB.GetFeedsDesc(context.Background(), database.GetFeedsDescParams{
		Limit:  1,
		Offset: 0,
	})
	if err == nil {
		t.Errorf("Expected 'error' but found 'nil'")
	}
	fmt.Println("Test Case 2 - Invalid database")
	dbconfig, err = InitConn("postgres://postgres:Re_Ghoul@127.0.0.1:5432/test_db?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	_, err = dbconfig.DB.GetFeedsDesc(context.Background(), database.GetFeedsDescParams{
		Limit:  1,
		Offset: 0,
	})
	if err == nil {
		t.Errorf("Expected 'error' but found 'nil'")
	}
	fmt.Println("Test Case 3 - Valid connection url")
	dbconfig, err = InitConn("postgres://postgres:Re_Ghoul@127.0.0.1:5432/" + utils.LoadEnv("db_database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	_, err = dbconfig.DB.GetFeedsDesc(context.Background(), database.GetFeedsDescParams{
		Limit:  1,
		Offset: 0,
	})
	if err != nil {
		t.Errorf("Expected 'nil' but found 'error'")
	}
}

func TestInitConfigWoo(t *testing.T) {
	fmt.Println("Test Case 1 - ")
}
