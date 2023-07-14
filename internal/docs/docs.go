package docs

import (
	"objects"
	"time"
)

// return all endpoints
func Endpoints() objects.Endpoints {
	return objects.Endpoints{
		Status:      true,
		Description: "Blog Aggregator API Documentation",
		Routes:      createRoutes(),
		Version:     "v1",
		Time:        time.Now().UTC(),
	}
}

func createRoutes() map[string]objects.Route {
	routes := map[string]objects.Route{
		"GET /v1/readiness": {
			Description:   "Returns the status of the API",
			Supports:      []string{"GET"},
			Params:        map[string]objects.Params{},
			AcceptsData:   false,
			Format:        []string{},
			Authorization: "None",
		},
		"GET /v1/err": {
			Description:   "Returns an internal server error",
			Supports:      []string{"GET"},
			Params:        map[string]objects.Params{},
			AcceptsData:   false,
			Format:        []string{},
			Authorization: "None",
		},
		"GET /v1/users": {
			Description:   "Returns user information",
			Supports:      []string{"GET"},
			Params:        map[string]objects.Params{},
			AcceptsData:   false,
			Format:        []string{},
			Authorization: "Authorization: ApiKey <key>",
		},
		"POST /v1/users": {
			Description:   "Creates a new user",
			Supports:      []string{"POST"},
			Params:        map[string]objects.Params{},
			AcceptsData:   true,
			Format:        objects.RequestBodyUser{},
			Authorization: "None",
		},
		"POST /v1/feeds": {
			Description:   "Creates a new feed",
			Supports:      []string{"POST"},
			Params:        map[string]objects.Params{},
			AcceptsData:   true,
			Format:        objects.RequestBodyFeed{},
			Authorization: "Authorization: ApiKey <key>",
		},
	}
	return routes
}
