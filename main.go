package main

import (
	"httprouter"
	"log"
	"net/http"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(httprouter.MiddleWare()))

	v1 := chi.NewRouter()

	v1.Get("/readiness", httprouter.ReadiHandle)
	v1.Get("/err", httprouter.ErrHandle)
	r.Mount("/v1", v1)

	port := utils.LoadEnv("port")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
