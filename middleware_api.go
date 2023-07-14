package main

import (
	"blog/internal/database"
	"net/http"
	"utils"
)

// custom Authhandler
type authHandler func(w http.ResponseWriter, r *http.Request, dbuser database.User)

// Authentication middleware
func (dbConfig *dbConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.ExtractAPIKey(r.Header.Get("Authorization"))
		if apiKey == "" {
			RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		dbUser, err := dbConfig.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		handler(w, r, dbUser)
	}
}
