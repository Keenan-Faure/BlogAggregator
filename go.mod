module blog

replace internal/utils => ./internal/utils

replace internal/httprouter => ./internal/httprouter

replace internal/objects => ./internal/objects

replace internal/docs => ./internal/docs

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/go-chi/cors v1.2.1
)
