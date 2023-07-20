package objects

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CreateFeedParam struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CreateUserParam struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
}

type CreatePostParams struct {
	Title       string         `json:"title"`
	Url         string         `json:"url"`
	Description sql.NullString `json:"description"`
	PublishedAt sql.NullTime   `json:"published_at"`
}

type ResponseBodyFeed struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Url           string    `json:"url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        uuid.UUID `json:"user_id"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}

type ResponseBodyPosts struct {
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

type RequestBodyUser struct {
	Name string `json:"name"`
}

type RequestBodyFeed struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RequestBodyFeedFollow struct {
	FeedID string `json:"feed_id"`
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

// Fetch RSS Feed structure
type RSS struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}
