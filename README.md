# BlogAggregator

A Blog Aggregator that fetches data from remote sources

## Main Features

- Building a RESTful API in Go
- Uses production-ready database tools like [PostgreSQL](https://www.postgresql.org), [SQLc](https://sqlc.dev), [Goose](https://github.com/pressly/goose), and [pgAdmin](https://www.pgadmin.org)
- Database migration using goose
- Use of Googles [UUID](https://pkg.go.dev/github.com/google/uuid) package
- Fetching and encoding of xml RSS feeds
- Use of [Wait Groups](https://pkg.go.dev/sync#WaitGroup) to bulk process RSS feeds

## Prerequistes

### Install POSTGRESQL and check if install was successful

Mac OS:

```
brew install postgresql

brew services start postgresql
```

Linux (WSL):

```
sudo apt update
sudo apt install postgresql postgresql-contrib

sudo service postgresql start
```

### Install sqlc

```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc version
```

### Install Goose

```
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -version
```

### Install the Go PostgreSQL driver

This module is used by `database/sql` and the `_` when importing indicates this.

```
go get github.com/lib/pq
```

### Configuring the .env file

By default you should have an `.env` file in your directory. However, these are default values!
Kindly populate them with the correct data.

## How to run the migrations:

- Navigate to the `sql/schema` directory and run `goose postgres {{CONN}} up`

Where `CONN` is the connection string for your database which follow the format:

```
protocol://username:password@host:port/database
```

## API Documentation

| Endpoint                          | Description                       | HTTP Method | Authorization | Params                         | Format                                       |
| --------------------------------- | :-------------------------------- | ----------- | ------------- | ------------------------------ | -------------------------------------------- |
| `/v1/readiness`                   | Returns the status of the API     | GET         | N/A           |                                |                                              |
| `/v1/err`                         | Returns an internal server error  | GET         | N/A           |                                |                                              |
| `/v1/users`                       | Returns user information          | GET         | ApiKey <key>  |                                |                                              |
| `/v1/users`                       | Creates a new user                | POST        | N/A           |                                | `json{"name": "UserName"}`                   |
| `/v1/feeds`                       | Gets all feeds                    | GET         | N/A           | `page={pageNum}&sort={method}` |                                              |
| `/v1/feeds`                       | Creates a new feed                | POST        | ApiKey <key>  |                                | `json{"name": "FeedName", "url": "FeedURL"}` |
| `/v1/feed_follows`                | Follows a feed                    | POST        | ApiKey <key>  |                                | `json{"feed_id": "FeedID"}`                  |
| `/v1/feed_follows/{feedFollowID}` | Unfollows a feed                  | DELETE      | N/A           |                                |                                              |
| `/v1/posts`                       | Displays posts followed by a user | GET         | ApiKey <key>  | `page={pageNum}&sort={method}` |                                              |

## Extensions (2 weeks extra)

- Support pagination of the endpoints that can return many items
  - Use query params to allow the user to select which page they wish to see
  - `?page=1`
  - Each page can have a limit of 10 `LIMIT $1 OFFSET $(PageNumber * 10)`
- Support different options for sorting and filtering posts using query parameters
  - Add a search that searches for the post/feed (title)
- Add support for other types of feeds (e.g. Atom, JSON, etc.) (WooCommerce)
  - JSON product data when authenticated (WooCommerce) can scrape product data to add to the website
- Add integration tests that use the API to create, read, update, and delete feeds and posts
  - Will be used in the webUI
- Create a simple web UI that uses your backend API
  - Landing page containing Navbar (Products, Feeds, Posts, Followed, Liked, bookmarked)
  - Little CSS mostly HTML
- Add bookmarking or "liking" to posts
  - Add extra table that marks which user liked or book marked feeds/posts
