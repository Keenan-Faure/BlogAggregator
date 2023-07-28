// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    ID, created_at, updated_at, name, api_key
) 
VALUES (
    $1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING id, created_at, updated_at, name, api_key
`

type CreateUserParams struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const deleteTestUsers = `-- name: DeleteTestUsers :exec

DELETE FROM users
WHERE id = $1
`

// >> used for tests << --
func (q *Queries) DeleteTestUsers(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTestUsers, id)
	return err
}

const getUser = `-- name: GetUser :one
select id, created_at, updated_at, name, api_key from users
WHERE "name" = $1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPI = `-- name: GetUserByAPI :one
SELECT id, created_at, updated_at, name, api_key FROM users
WHERE api_key = $1
`

func (q *Queries) GetUserByAPI(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPI, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
