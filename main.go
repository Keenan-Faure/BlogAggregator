package main

import (
	"log"
	"net/http"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(MiddleWare()))

	v1 := chi.NewRouter()

	v1.Get("/readiness", ReadiHandle)
	v1.Get("/err", ErrHandle)
	v1.Get("/users", GetUserByApiKeyHandle)
	v1.Post("/users", CreateUserHandle)
	r.Mount("/v1", v1)

	port := utils.LoadEnv("port")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
