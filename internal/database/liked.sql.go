// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: liked.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createLike = `-- name: CreateLike :one
INSERT INTO liked (
    id,
    post_id,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, post_id, user_id, created_at, updated_at
`

type CreateLikeParams struct {
	ID        uuid.UUID `json:"id"`
	PostID    uuid.UUID `json:"post_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (Liked, error) {
	row := q.db.QueryRowContext(ctx, createLike,
		arg.ID,
		arg.PostID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Liked
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLikedByPostID = `-- name: GetLikedByPostID :one
SELECT id, post_id, user_id, created_at, updated_at FROM liked
WHERE post_id = $1
`

func (q *Queries) GetLikedByPostID(ctx context.Context, postID uuid.UUID) (Liked, error) {
	row := q.db.QueryRowContext(ctx, getLikedByPostID, postID)
	var i Liked
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const removeLikedByPostID = `-- name: RemoveLikedByPostID :one
DELETE FROM liked
WHERE post_id = $1
RETURNING id, post_id, user_id, created_at, updated_at
`

func (q *Queries) RemoveLikedByPostID(ctx context.Context, postID uuid.UUID) (Liked, error) {
	row := q.db.QueryRowContext(ctx, removeLikedByPostID, postID)
	var i Liked
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const removeLikedByUserID = `-- name: RemoveLikedByUserID :one
DELETE FROM liked
WHERE user_id = $1
RETURNING id, post_id, user_id, created_at, updated_at
`

func (q *Queries) RemoveLikedByUserID(ctx context.Context, userID uuid.UUID) (Liked, error) {
	row := q.db.QueryRowContext(ctx, removeLikedByUserID, userID)
	var i Liked
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const searchLikedByPostTitle = `-- name: SearchLikedByPostTitle :many
SELECT liked.id, liked.post_id, liked.user_id, liked.created_at, liked.updated_at FROM liked
INNER JOIN posts
ON liked.post_id = posts.id
WHERE posts.title = $1
`

func (q *Queries) SearchLikedByPostTitle(ctx context.Context, title string) ([]Liked, error) {
	rows, err := q.db.QueryContext(ctx, searchLikedByPostTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Liked
	for rows.Next() {
		var i Liked
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}