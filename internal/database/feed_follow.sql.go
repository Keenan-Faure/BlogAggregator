// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows(
    id,
    feed_id,
    user_id,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, feed_id, user_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.FeedID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedByID = `-- name: DeleteFeedByID :one
DELETE FROM feed_follows WHERE id = $1
RETURNING id, feed_id, user_id, created_at, updated_at
`

func (q *Queries) DeleteFeedByID(ctx context.Context, id uuid.UUID) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, deleteFeedByID, id)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedFollow = `-- name: GetFeedFollow :one
SELECT id, feed_id, user_id, created_at, updated_at FROM feed_follows
WHERE feed_id = $1 AND user_id = $2
`

type GetFeedFollowParams struct {
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) GetFeedFollow(ctx context.Context, arg GetFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, getFeedFollow, arg.FeedID, arg.UserID)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, feed_id, user_id, created_at, updated_at FROM feed_follows
`

func (q *Queries) GetFeedFollows(ctx context.Context) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
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

const getUserFeedFollows = `-- name: GetUserFeedFollows :many
SELECT id, feed_id, user_id, created_at, updated_at FROM feed_follows
WHERE user_id = $1
`

func (q *Queries) GetUserFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getUserFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
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
