package main

import (
	"blog/internal/database"
	"database/sql"
	"log"
)

// Initiates a connection to the database and
// if successful returns the connection
func InitConn(dbURL string) (dbConfig, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
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
