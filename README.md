# BlogAggregator

A Blog Aggregator that fetches data from remote sources

## Main Features

-   Building a RESTful API in Go, and you'll
-   Uses production-ready database tools like [PostgreSQL](https://www.postgresql.org), (SQLc)[https://sqlc.dev], [Goose](https://github.com/pressly/goose), and [pgAdmin](https://www.pgadmin.org)
-   Database migration
-   Use of Googles [UUID](https://pkg.go.dev/github.com/google/uuid) package

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

-   Navigate to the `sql/schema` directory and run `goose postgres {{CONN}} up`

Where `CONN` is the connection string for your database which follow the format:

```
protocol://username:password@host:port/database
```
