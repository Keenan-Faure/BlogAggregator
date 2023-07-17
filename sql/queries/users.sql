-- name: CreateUser :one
INSERT INTO users(
    ID, created_at, updated_at, name, api_key
) 
VALUES (
    $1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPI :one
SELECT * FROM users
WHERE api_key = $1;

-- name: GetUser :one
select * from users
WHERE "name" = $1;

-- https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html