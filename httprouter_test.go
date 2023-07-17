package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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

func GetURL(url string) (*http.Response, error) {
	return http.Get(url)
}

func PostURL(url string, bodyData map[string]any) (*http.Response, error) {
	jsonStr, err := json.Marshal(bodyData)
	if err != nil {
		return &http.Response{}, err
	}
	return http.Post(url, "application/json", bytes.NewReader(jsonStr))
}
