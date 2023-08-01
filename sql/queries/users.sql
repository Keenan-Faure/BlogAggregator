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

-- name: GetUserByName :one
select * from users
WHERE "name" = $1;

-- >> used for tests << --

-- name: DeleteTestUsers :exec
DELETE FROM users
WHERE id = $1;