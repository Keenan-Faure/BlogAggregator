package main

import (
	"blog/internal/database"
	"context"
	"database/sql"
	"log"
)

// Initiates a connection to the database and
// if successful returns the connection
func InitConn(dbURL string) (dbConfig, error) {
	log.Println(dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return dbConfig{}, err
	}
	return storeConfig(db), nil
}

// Stores the database connection inside a config struct
func storeConfig(conn *sql.DB) dbConfig {
	_, err := database.New(conn).GetFeedsDesc(context.Background(), database.GetFeedsDescParams{
		Limit:  1,
		Offset: 0,
	})
	log.Println(err)
	if err == nil {
		config := dbConfig{
			DB:    database.New(conn),
			Valid: true,
		}
		return config
	} else {
		config := dbConfig{
			DB:    database.New(conn),
			Valid: false,
		}
		return config
	}
}
