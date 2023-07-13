package main

import (
	"blog/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"objects"
	"time"

	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

// v1 handlers

// test for RespondWithJSON
func ReadiHandle(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, objects.ReadyHandle{
		Status: "ok",
	})
}

// error handler
func ErrHandle(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 200, "Internal Server Error")
}

func CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	config, err := InitConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	decoder := json.NewDecoder(r.Body)
	params := objects.RequestBodyUser{}
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	dbUser, err := config.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, dbUser)
}

// middleware that determines which headers, http methods and orgins are allowed
func MiddleWare() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}
}

// helper function
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

// helper function
func RespondWithError(w http.ResponseWriter, code int, msg string) error {
	return RespondWithJSON(w, code, map[string]string{"error": msg})
}
