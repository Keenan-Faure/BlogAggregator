package objects

import (
	"time"
)

type CreateUserParam struct {
	ID        string `json:"id"`
	CreateAt  string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
}

type RequestBodyUser struct {
	Name string `json:"name"`
}

type NoResponse struct {
}

type ResponseBodyUser struct {
	ID        string `json:"id"`
	CreateAt  string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

type ReadyHandle struct {
	Status string `json:"status"`
}

type ErrorHandle struct {
	Error string `json:"error"`
}

// Docs Endpoints
type Endpoints struct {
	Status      bool             `json:"status"`
	Description string           `json:"description"`
	Routes      map[string]Route `json:"routes"`
	Version     string           `json:"version"`
	Time        time.Time        `json:"time"`
}

type Route struct {
	Description   string            `json:"description"`
	Supports      []string          `json:"supports"`
	Params        map[string]Params `json:"params"`
	AcceptsData   bool              `json:"accepts_data"`
	Format        any               `json:"format"`
	Authorization string            `json:"auth"`
}

type Params struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
