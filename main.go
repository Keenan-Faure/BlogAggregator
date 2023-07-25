package main

import (
	"blog/internal/database"
	"log"
	"net/http"
	"productfetch"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DB *database.Queries
}

func main() {
	dbconfig, err := InitConn(utils.LoadEnv("db_url"))
	if err != nil {
		log.Fatal(err.Error())
	}
	woo, err := productfetch.InitConfigWoo()
	if err != nil {
		log.Println(err)
	}
	shopify, err := productfetch.InitConfigShopify()
	if err != nil {
		log.Println(err)
	}

	FetchWorker(dbconfig, shopify, woo)

	r := chi.NewRouter()
	r.Use(cors.Handler(MiddleWare()))
	v1 := chi.NewRouter()

	v1.Get("/readiness", ReadiHandle)
	v1.Get("/err", ErrHandle)
	v1.Get("/users", dbconfig.middlewareAuth(dbconfig.GetUserByApiKeyHandle))
	v1.Post("/users", dbconfig.CreateUserHandle)
	v1.Post("/feed", dbconfig.middlewareAuth(dbconfig.CreateFeedHandler))
	v1.Get("/feed", dbconfig.GetAllFeedsHandle)
	v1.Get("/feed_search", dbconfig.SearchFeedHandle)
	v1.Get("/feed_follows", dbconfig.middlewareAuth(dbconfig.GetFeedFollowHandle))
	v1.Post("/feed_follows", dbconfig.middlewareAuth(dbconfig.CreateFeedFollowHandle))
	v1.Delete("/feed_follows/{feedFollowID}", dbconfig.UnFollowFeedHandle)
	v1.Get("/posts", dbconfig.middlewareAuth(dbconfig.GetPostFollowHandle))
	v1.Get("/posts_search", dbconfig.SearchPostHandle)

	v1.Get("/", dbconfig.Endpoints)
	r.Mount("/v1", v1)

	port := utils.LoadEnv("port")
	if port == "" {
		log.Fatal("Port not defined in Environment")
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
