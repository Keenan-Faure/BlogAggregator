#!/bin/bash
cd /web-app/sql/schema
source .env

echo "Checking GOOSE version"
goose -version

SSL_MODE="?sslmode=disable"
DB_STRING="${DB_URL}${DATABASE}${SSL_MODE}"

echo $DB_STRING

echo "running migrations on '${DATABASE}'"

goose postgres "$DB_STRING" up