module blog

replace internal/utils => ./internal/utils

replace internal/objects => ./internal/objects

replace internal/docs => ./internal/docs

replace internal/dbconfig => ./internal/dbconfig

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/go-chi/cors v1.2.1
)

require (
	github.com/google/uuid v1.3.0 // direct
	github.com/lib/pq v1.10.9 // direct
)
