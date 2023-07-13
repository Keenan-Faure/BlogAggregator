package main

import (
	"blog/internal/database"
	"database/sql"
	"utils"
)

type dbConfig struct {
	DB *database.Queries
}

// Initiates a connection to the database and
// if successful returns the connection
func InitConn() (dbConfig, error) {
	dbURL := utils.LoadEnv("db_url")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return dbConfig{}, err
	}
	return storeConfig(db), nil
}

// Stores the database connection inside a config struct
func storeConfig(conn *sql.DB) dbConfig {
	config := dbConfig{
		DB: database.New(conn),
	}
	return config
}
