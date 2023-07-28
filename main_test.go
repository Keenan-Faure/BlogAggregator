package main

import (
	"blog/internal/database"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"objects"
	"productfetch"
	"strconv"
	"time"
	"utils"

	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// Tests below assume that the tables exist
// and the correct database was chosen
// all data created by tests are removed upon finish

func UFetchHelperPost(endpoint, method string, auth string, body io.Reader) (*http.Response, error) {
	httpClient := http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest(method, "http://localhost:"+utils.LoadEnv("port")+"/v1/"+endpoint, body)
	if auth != "" {
		req.Header.Add("Authorization", "ApiKey "+auth)
	}
	if err != nil {
		log.Println(err)
		return &http.Response{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return &http.Response{}, err
	}
	return res, nil
}

func UCreateFeed(ApiKey string) database.Feed {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(objects.RequestBodyFeed{
		Name: "test_feed_xyz_123_456",
		URL:  "no_one.would_have_this_name.com",
	})
	res, _ := UFetchHelperPost("feed", "POST", ApiKey, &buffer)
	defer res.Body.Close()
	respBody, _ := io.ReadAll(res.Body)
	type FeedCreation struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}
	feedData := FeedCreation{}
	json.Unmarshal(respBody, &feedData)
	return feedData.Feed
}

func UCreateUser() database.User {
	userBody := database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      "test_abc123_xyz_def",
	}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(userBody)
	res, _ := UFetchHelperPost("users", "POST", "", &buffer)
	defer res.Body.Close()
	respBody, _ := io.ReadAll(res.Body)
	userData := database.User{}
	json.Unmarshal(respBody, &userData)
	return userData
}

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
	api_key = "ck_21233ljosidjalksd"
	api_secret = "cs_123123nasldnasd"
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
	time.Sleep(time.Second * 5)
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

func TestFetchWorkerShopify(t *testing.T) {
	fmt.Println("Test Case 1 - using shopify api url in env")

	store_name_shopify := utils.LoadEnv("t_store_name")
	api_key_shopify := utils.LoadEnv("t_api_key")
	api_password_shopify := utils.LoadEnv("t_api_password")
	version := utils.LoadEnv("t_version")
	shopify := productfetch.InitConfigShopify(store_name_shopify, api_key_shopify, api_password_shopify, version)
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}

	time_now := time.Now().UTC()
	FetchWorker(dbconfig, shopify, productfetch.ConfigWoo{})
	time.Sleep(time.Second * 5)

	product, err := dbconfig.DB.GetFirstRecordShopify(context.Background())
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	if time_now.Compare(product.UpdatedAt) != -1 {
		t.Errorf("Expected 'UpdatedAt' to be after 'time_now'")
	}
	dbconfig.DB.DeleteTestShopifyProducts(context.Background(), store_name_shopify)
}

func TestFetchWorkerWoo(t *testing.T) {
	fmt.Println("Test Case 1 - using woocommerce api url in env")

	store_name := utils.LoadEnv("t_woo_store_name")
	api_key := utils.LoadEnv("t_woo_consumer_key")
	api_secret := utils.LoadEnv("t_woo_consumer_secret")
	woo := productfetch.InitConfigWoo(store_name, api_key, api_secret)
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}

	time_now := time.Now().UTC()
	FetchWorker(dbconfig, productfetch.ConfigShopify{}, woo)
	time.Sleep(time.Second * 5)

	product, err := dbconfig.DB.GetFirstRecordWoo(context.Background())
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	if time_now.Compare(product.UpdatedAt) != -1 {
		t.Errorf("Expected 'UpdatedAt' to be after 'time_now'")
	}
	dbconfig.DB.DeleteTestWooProducts(context.Background(), store_name)
}

func UFetchHelper(endpoint, method, auth string) (*http.Response, error) {
	httpClient := http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest(method, "http://localhost:"+utils.LoadEnv("port")+"/v1/"+endpoint, nil)
	if auth != "" {
		req.Header.Add("Authorization", "ApiKey "+auth)
	}
	if err != nil {
		log.Println(err)
		return &http.Response{}, err
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return &http.Response{}, err
	}
	return res, nil
}

func TestErrEndpoint(t *testing.T) {
	fmt.Println("Test 1 - testing GET /err")
	type ErrorStruct struct {
		Error string `json:"error"`
	}
	res, err := UFetchHelper("err", "GET", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	data := ErrorStruct{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if data.Error != "Internal Server Error" {
		t.Errorf("Expected 'Internal Server Error' but found: " + data.Error)
	}

	fmt.Println("Test 2 - testing POST /err")
	res, err = UFetchHelper("err", "POST", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 405 {
		t.Errorf("Expected '405' but found: " + strconv.Itoa(res.StatusCode))
	}
}

func TestReadinessEndpoint(t *testing.T) {
	fmt.Println("Test 1 - testing GET /readiness")
	type readinessStruct struct {
		Status string `json:"status"`
	}
	res, err := UFetchHelper("readiness", "GET", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	data := readinessStruct{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if data.Status != "ok" {
		t.Errorf("Expected 'ok' but found: " + data.Status)
	}

	fmt.Println("Test 2 - testing POST /readiness")
	res, err = UFetchHelper("err", "POST", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 405 {
		t.Errorf("Expected '405' but found: " + strconv.Itoa(res.StatusCode))
	}
}

func TestUserCrud(t *testing.T) {
	fmt.Println("Test 1 - Creating user")
	userBody := database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      "test_abc123_xyz_def",
	}
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(userBody)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	res, err := UFetchHelperPost("users", "POST", "", &buffer)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 201 {
		t.Errorf("Expected '201' but found: " + strconv.Itoa(res.StatusCode))
	}
	userData := database.User{}
	err = json.Unmarshal(respBody, &userData)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if userData.Name != "test_abc123_xyz_def" {
		t.Errorf("Expected 'test_abc123_xyz_def' but found: " + userData.Name)
	}

	fmt.Println("Test 2 - Fetching user")
	res, err = UFetchHelperPost("users", "GET", userData.ApiKey, nil)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 202 {
		t.Errorf("Expected '202' but found: " + strconv.Itoa(res.StatusCode))
	}
	userData = database.User{}
	err = json.Unmarshal(respBody, &userData)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if userData.Name != "test_abc123_xyz_def" {
		t.Errorf("Expected 'test_abc123_xyz_def' but found: " + userData.Name)
	}

	fmt.Println("Test 3 - Deleting user & recheck")
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	dbconfig.DB.DeleteTestUsers(context.Background(), userData.ID)
	type ErrorStruct struct {
		Error string `json:"error"`
	}
	res, err = UFetchHelper("users", "GET", userData.ApiKey)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 404 {
		t.Errorf("Expected '404' but found: " + strconv.Itoa(res.StatusCode))
	}
	data := ErrorStruct{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if data.Error != "record not found" {
		t.Errorf("Expected 'record not found' but found: " + data.Error)
	}
}

func TestFeedCrud(t *testing.T) {
	fmt.Println("Test 1 - Creating Feed")
	user := UCreateUser()
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(objects.RequestBodyFeed{
		Name: "test_feed_xyz_123_456",
		URL:  "no_one.would_have_this_name.com",
	})
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	res, err := UFetchHelperPost("feed", "POST", user.ApiKey, &buffer)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 201 {
		t.Errorf("Expected '201' but found: " + strconv.Itoa(res.StatusCode))
	}
	type FeedCreation struct {
		Feed       objects.ResponseBodyFeed `json:"feed"`
		FeedFollow database.FeedFollow      `json:"feed_follow"`
	}
	feedData := FeedCreation{}
	err = json.Unmarshal(respBody, &feedData)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if feedData.Feed.Name != "test_feed_xyz_123_456" {
		t.Errorf("Expected 'test_feed_xyz_123_456' but found: " + feedData.Feed.Name)
	}

	fmt.Println("Test 2 -  Fetching Feed")
	res, err = UFetchHelper("feed_search?q=test_feed_xyz_123_456", "GET", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	feeds := []database.Feed{}
	err = json.Unmarshal(respBody, &feeds)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if feeds[0].Name != "test_feed_xyz_123_456" {
		t.Errorf("Expected 'test_feed_xyz_123_456' but found: " + feeds[0].Name)
	}

	fmt.Println("Test 3 - Deleting Feed & recheck")
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	dbconfig.DB.DeleteTestFeeds(context.Background(), user.ID)
	res, err = UFetchHelper("feed_search?q=test_feed_xyz_123_456", "GET", "")
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	feeds = []database.Feed{}
	err = json.Unmarshal(respBody, &feeds)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if len(feeds) > 0 {
		if feeds[0].Name != "test_feed_xyz_123_456" {
			t.Errorf("Expected 'test_feed_xyz_123_456' but found: " + feeds[0].Name)
		}
	}
	dbconfig.DB.DeleteTestUsers(context.Background(), user.ID)
}

func TestFeedFollowsCrud(t *testing.T) {
	fmt.Println("Test 1 - Creating & Fetching of Feed-follows")
	user := UCreateUser()
	feed := UCreateFeed(user.ApiKey)
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(objects.RequestBodyFeedFollow{
		FeedID: feed.ID.String(),
	})
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	res, err := UFetchHelper("feed_follows", "GET", user.ApiKey)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	feedsFollows := []database.FeedFollow{}
	err = json.Unmarshal(respBody, &feedsFollows)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if feedsFollows[0].FeedID.String() != feed.ID.String() {
		t.Errorf("Expected '" + feed.ID.String() + "' but found: " + feedsFollows[0].FeedID.String())
	}

	fmt.Println("Test 2 - Deleting Feed-follows")
	res, err = UFetchHelper("feed_follows/"+feedsFollows[0].ID.String(), "DELETE", user.ApiKey)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	defer res.Body.Close()
	respBody, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected '200' but found: " + strconv.Itoa(res.StatusCode))
	}
	feedsFollow := database.FeedFollow{}
	err = json.Unmarshal(respBody, &feedsFollow)
	if err != nil {
		t.Errorf("expected 'nil' but found: " + err.Error())
	}
	if feedsFollow.FeedID.String() != feed.ID.String() {
		t.Errorf("Expected '" + feed.ID.String() + "' but found: " + feedsFollow.FeedID.String())
	}
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		t.Errorf("Expected 'nil' but found: " + err.Error())
	}
	dbconfig.DB.DeleteTestFeeds(context.Background(), user.ID)
	dbconfig.DB.DeleteTestUsers(context.Background(), user.ID)
}
