package main

import (
	"blog/internal/database"
	"flag"
	"log"
	"net/http"
	"productfetch"
	"utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DB    *database.Queries
	Valid bool
}

func main() {
	dbconfig, err := InitConn(utils.LoadEnv("db_url") + utils.LoadEnv("database") + "?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	dbg := flag.Bool("test", false, "Enable server for tests only")
	flag.Parse()

	if !*dbg {
		store_name := utils.LoadEnv("woo_store_name")
		api_key := utils.LoadEnv("woo_consumer_key")
		api_secret := utils.LoadEnv("woo_consumer_secret")
		woo := productfetch.InitConfigWoo(store_name, api_key, api_secret)

		store_name_shopify := utils.LoadEnv("store_name")
		api_key_shopify := utils.LoadEnv("api_key")
		api_password_shopify := utils.LoadEnv("api_password")
		version := utils.LoadEnv("version")
		shopify := productfetch.InitConfigShopify(store_name_shopify, api_key_shopify, api_password_shopify, version)

		FetchWorker(dbconfig, shopify, woo)
	}
	setUpAPI(dbconfig)
}

// starts up the API
func setUpAPI(dbconfig dbConfig) {
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
	v1.Delete("/feed_follows/{feedFollowID}", dbconfig.middlewareAuth(dbconfig.UnFollowFeedHandle))
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
