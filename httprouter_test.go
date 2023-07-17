package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Init_Main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(MiddleWare()))

	v1 := chi.NewRouter()

	dbconfig, err := InitConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	v1.Get("/readiness", ReadiHandle)
	v1.Get("/err", ErrHandle)
	v1.Get("/users", dbconfig.middlewareAuth(dbconfig.GetUserByApiKeyHandle))
	v1.Post("/users", dbconfig.CreateUserHandle)
	v1.Post("/feed", dbconfig.middlewareAuth(dbconfig.CreateFeedHandler))
	v1.Get("/feed", dbconfig.GetAllFeedsHandle)
	v1.Get("/feed_follows", dbconfig.middlewareAuth(dbconfig.GetFeedFollowHandle))
	v1.Post("/feed_follows", dbconfig.middlewareAuth(dbconfig.CreateFeedFollowHandle))
	v1.Delete("/feed_follows/{feedFollowID}", dbconfig.UnFollowFeedHandle)
	r.Mount("/v1", v1)

	port := utils.LoadEnv("port")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}

func TestCheckFeedFollowExist(t *testing.T) {
	Init_Main()
	fmt.Println("Test case 1 - valid feed id")

}

//helper functions

func GetURL(url string, bodyData map[string]any) (*http.Response, error) {
	client := &http.Client{}
	req, err := GetRequest(url, "GET", bodyData)
	if err != nil {
		return &http.Response{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func PostURL(url string, bodyData map[string]any) (*http.Response, error) {
	client := &http.Client{}
	req, err := GetRequest(url, "POST", bodyData)
	if err != nil {
		return &http.Response{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func DeleteURL(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := GetRequest(url, "DELETE", nil)
	if err != nil {
		return &http.Response{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func GetRequest(url string, option string, bodyData map[string]any) (*http.Request, error) {
	if bodyData == nil {
		bodyData = make(map[string]any)
	}
	jsonStr, err := json.Marshal(bodyData)
	if err != nil {
		return &http.Request{}, err
	}
	req, err := http.NewRequest(option, url, bytes.NewReader(jsonStr))
	if err != nil {
		return &http.Request{}, err
	}
	return req, err
}
