# BlogAggregator

A Blog Aggregator that fetches data from remote sources with added extensions.

Project can be found on [Boot.Dev](https://boot.dev)

## Main Features

- A RESTful API in Go
- Uses production-ready database tools like [PostgreSQL](https://www.postgresql.org), [SQLc](https://sqlc.dev), [Goose](https://github.com/pressly/goose), and [pgAdmin](https://www.pgadmin.org)
- Database migration using goose
- Use of Googles [UUID](https://pkg.go.dev/github.com/google/uuid) package to uniquely identify a record in a database (ID).
- Fetching and encoding of xml RSS feeds
- Use of [Wait Groups](https://pkg.go.dev/sync#WaitGroup) to bulk process RSS feeds

### Extensions added

- Supports pagination of the endpoints that can return many items. Visit docs in `README.md` for more information.
- Fetching and storing of JSON product data from E-Commerce websites - Shopify and WooCommerce.
- Support different options for sorting and filtering posts using query parameters
- Added Bookmarking & Liking to posts.
- Has integration tests that use the API to create, read, update, and delete feeds, posts, bookmarks and likes
- Has a simple web UI that uses the backend API to Create, View, Follow Bookmark and Like feeds and posts.
- Add a Docker implementation of the project, more info found [below](https://github.com/Keenan-Faure/BlogAggregator#docker-implementation)

## Prerequistes

This assumes the code already exists on your local machine, hence, please download the project code files.

### 1. Install POSTGRESQL and check if install was successful

Mac OS:

```bash
brew install postgresql

brew services start postgresql
```

Linux (WSL):

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib

sudo service postgresql start
```

### 2. Install sqlc

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc version
```

### 3. Install Goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -version
```

### 4. Install the Go PostgreSQL driver

This module is used by `database/sql` and the `_` when importing indicates this.

```bash
go get github.com/lib/pq
```

### 5. Configuring the .env file

By default you should have an `.env` file in your directory. However, these are default values!
Populate them with the correct values as per your configuration.

### 6. Run the migrations

- Navigate to the `sql/schema` directory of the project and run `goose postgres {{CONN}} up`

Where `CONN` is the connection string for your database which follow the format:

```bash
protocol://username:password@host:port/database
```

Note that if the prerequistes are not completed, then the application will fail to run (ಥ \_ ಥ)

## Docker implementation

Note the docker container on DockerHub is not available yet :o
Hence, I have left out that implementation.

This assumes the project file already exist locally on your machine and that you have navigated to the local directory
in your command line (`bash`, `zsh`)

### Installing Docker

Depending on your Operating System download and install [Docker](https://www.docker.com/products/docker-desktop/)

### Running the shell commands

Outline of shell commands

| Command      | Description                                                                      |
| ------------ | -------------------------------------------------------------------------------- |
| `./run.sh`   | Starts the docker containers and the application                                 |
| `./stop.sh`  | Stops the docker containers                                                      |
| `./reset.sh` | Removes all containers, volumes & images locally (does not run `./run.sh` again) |

## API Integration Tests

To run these unit tests, please note that you would have to start the program using the `--test` param as shown below:

```bash
go build -o out && ./out --test
```

This will **only** start up the API server, after which you can proceed to running the tests found in `main_test.go`

## API Documentation

| Endpoint                          | Description                         | HTTP Method | Authorization | Params                         | Format                                       |
| --------------------------------- | :---------------------------------- | ----------- | ------------- | ------------------------------ | -------------------------------------------- |
| `/v1/readiness`                   | Returns the status of the API       | GET         |               |                                |                                              |
| `/v1/err`                         | Returns an internal server error    | GET         |               |                                |                                              |
| `/v1/users`                       | Returns user information            | GET         | ApiKey {key}  |                                |                                              |
| `/v1/liked`                       | Returns Liked posts for a user      | GET         | ApiKey {key}  | `page={pageNum}&sort={method}` |                                              |
| `/v1/bookmark`                    | Returns Bookmarked posts for a user | GET         | ApiKey {key}  | `page={pageNum}&sort={method}` |                                              |
| `/v1/feeds`                       | Gets all feeds                      | GET         |               | `page={pageNum}&sort={method}` | `json{"name": "UserName"}`                   |
| `/v1/posts`                       | Displays posts followed by a user   | GET         | ApiKey {key}  | `page={pageNum}&sort={method}` |                                              |
| `/v1/users`                       | Creates a new user                  | POST        |               |                                | `json{"name": "UserName"}`                   |
| `/v1/bookmark`                    | Bookmarks a post                    | POST        | ApiKey {key}  |                                | `json{"post_id": "PostID"}`                  |
| `/v1/liked`                       | Likes a post                        | POST        | ApiKey {key}  |                                | `json{"post_id": "PostID"}`                  |
| `/v1/feeds`                       | Creates a new feed                  | POST        | ApiKey {key}  |                                | `json{"name": "FeedName", "url": "FeedURL"}` |
| `/v1/feed_follows`                | Follows a feed                      | POST        | ApiKey {key}  |                                | `json{"feed_id": "FeedID"}`                  |
| `/v1/feed_search`                 | Searches for a feed by name         | POST        |               | `q={FeedName}`                 |                                              |
| `/v1/posts_search`                | Searches for a post by title        | POST        |               | `q={PostTitle}`                |                                              |
| `/v1/bookmark_search`             | Searches for a Bookmark by title    | POST        |               | `q={PostTitle}`                |                                              |
| `/v1/liked_search`                | Searches for a Liked Post by title  | POST        |               | `q={PostTitle}`                |                                              |
| `/v1/liked/{postID}`              | Unlikes a post                      | DELETE      |               |                                |                                              |
| `/v1/bookmark/{postID}`           | Removed the bookmark                | DELETE      |               |                                |                                              |
| `/v1/feed_follows/{feedFollowID}` | Unfollows a feed                    | DELETE      |               |                                |                                              |
