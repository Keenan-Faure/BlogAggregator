package main

import (
	"blog/internal/database"
	"log"
	"net/http"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DB *database.Queries
}

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(MiddleWare()))

	v1 := chi.NewRouter()

	dbconfig, err := InitConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	v1.Get("/readiness", ReadiHandle)
	v1.Get("/err", ErrHandle)
	v1.Get("/users", dbconfig.GetUserByApiKeyHandle)
	v1.Post("/users", dbconfig.CreateUserHandle)
	r.Mount("/v1", v1)

	port := utils.LoadEnv("port")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
