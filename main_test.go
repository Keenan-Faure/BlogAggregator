package main

import (
	"blog/internal/database"
	"context"
	"database/sql"
	"fmt"
	"productfetch"
	"time"
	"utils"

	"testing"

	"github.com/google/uuid"
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
	dbconfig, err = InitConn(utils.LoadEnv("db_url") + "fake_abc123" + "?sslmode=disable")
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
	dbconfig, err = InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
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
	fmt.Println("Test Case 1 - Invalid WooCommerce variables")
	store_name := "test.com"
	api_key := "abc123"
	api_secret := "xyz456"
	wooConfig := productfetch.InitConfigWoo(store_name, api_key, api_secret)
	if wooConfig.Valid {
		t.Errorf("Expected 'false' but found 'true'")
	}
	fmt.Println("Test Case 2 - Valid WooCommerce variables")
	store_name = "test.com"
	api_key = "cs_21233ljosidjalksd"
	api_secret = "ck_123123nasldnasd"
	wooConfig = productfetch.InitConfigWoo(store_name, api_key, api_secret)
	if !wooConfig.Valid {
		t.Errorf("Expected 'true' but found 'false'")
	}
}

func TestInitConfigShopify(t *testing.T) {
	fmt.Println("Test Case 1 - Invalid Shopify variables")
	store_name := "test.com"
	api_key := ""
	api_secret := "test_secret"
	version := "2023-07"
	wooConfig := productfetch.InitConfigShopify(store_name, api_key, api_secret, version)
	if wooConfig.Valid {
		t.Errorf("Expected 'false' but found 'true'")
	}
	fmt.Println("Test Case 2 - Valid Shopify variables")
	store_name = "test.com"
	api_key = "test_key"
	api_secret = "test_secret"
	version = "2023-07"
	wooConfig = productfetch.InitConfigShopify(store_name, api_key, api_secret, version)
	if !wooConfig.Valid {
		t.Errorf("Expected 'true' but found 'false'")
	}
}

func TestFetchWorker(t *testing.T) {
	fmt.Println("Test Case 1 - valid db connection")
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	time_now := time.Now().UTC()
	id_feed := uuid.New()
	id_user := uuid.New()
	_, err = dbconfig.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id_user,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      "Test User",
	})
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	_, err = dbconfig.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        id_feed,
		Name:      "Los Angeles Times",
		Url:       "https://www.latimes.com/local/rss2.0.xml",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UserID: id_user,
	})
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	FetchWorker(dbconfig, productfetch.ConfigShopify{}, productfetch.ConfigWoo{})
	time.Sleep(time.Second * 10)
	post, err := dbconfig.DB.GetFirstRecordPost(context.Background())
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	if time_now.Compare(post.UpdatedAt) != -1 {
		t.Errorf("Expected 'UpdatedAt' to be after 'time_now'")
	}
	dbconfig.DB.DeleteTestFeeds(context.Background(), id_user)
	dbconfig.DB.DeleteTestPosts(context.Background(), id_feed)
	dbconfig.DB.DeleteTestUsers(context.Background(), id_user)
}
