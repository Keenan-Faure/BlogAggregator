package main

import (
	"blog/internal/database"
	"encoding/json"
	"net/http"
	"objects"
	"time"

	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

// v1 handlers

func (dbConfig *dbConfig) CreateFeedHandler(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	params, err := DecodeFeedRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if FeedValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	feed, err := dbConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseFeedToObject(feed))
}

// returns a user with the specific ApiKey
func (dbConfig *dbConfig) GetUserByApiKeyHandle(w http.ResponseWriter, r *http.Request, dbUser database.User) {
	RespondWithJSON(w, http.StatusAccepted, DatabaseUserToObject(dbUser))
}

// creates a new user
func (dbConfig *dbConfig) CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	params, err := DecodeUserRequestBody(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if UserValidation(params) != nil {
		RespondWithError(w, http.StatusBadRequest, "data validation error")
		return
	}
	dbUser, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, DatabaseUserToObject(dbUser))
}

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
