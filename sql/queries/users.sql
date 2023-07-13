-- name: CreateUser :one
INSERT INTO users(
    ID, created_at, updated_at, name
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html