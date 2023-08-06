// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds(
    id,
    name,
    url,
    created_at,
    updated_at,
    last_fetched_at,
    user_id
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, url, created_at, updated_at, user_id, last_fetched_at
`

type CreateFeedParams struct {
	ID            uuid.UUID    `json:"id"`
	Name          string       `json:"name"`
	Url           string       `json:"url"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
	UserID        uuid.UUID    `json:"user_id"`
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.LastFetchedAt,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const deleteTestFeeds = `-- name: DeleteTestFeeds :exec
DELETE FROM feeds
WHERE user_id = $1
`

func (q *Queries) DeleteTestFeeds(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTestFeeds, userID)
	return err
}

const getFeed = `-- name: GetFeed :one
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
WHERE id = $1
`

func (q *Queries) GetFeed(ctx context.Context, id uuid.UUID) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
WHERE "url" = $1
`

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeedSearchName = `-- name: GetFeedSearchName :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds WHERE "name"
SIMILAR TO $1
`

func (q *Queries) GetFeedSearchName(ctx context.Context, similarToEscape string) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeedSearchName, similarToEscape)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.LastFetchedAt,
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

const getFeedsAsc = `-- name: GetFeedsAsc :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
ORDER BY updated_at ASC
LIMIT $1 OFFSET $2
`

type GetFeedsAscParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedsAsc(ctx context.Context, arg GetFeedsAscParams) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsAsc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.LastFetchedAt,
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

const getFeedsDesc = `-- name: GetFeedsDesc :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2
`

type GetFeedsDescParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeedsDesc(ctx context.Context, arg GetFeedsDescParams) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsDesc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.LastFetchedAt,
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

const getNextFeedsToFetch = `-- name: GetNextFeedsToFetch :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
ORDER BY last_fetched_at DESC
LIMIT $1
`

func (q *Queries) GetNextFeedsToFetch(ctx context.Context, limit int32) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getNextFeedsToFetch, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.LastFetchedAt,
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

const markFeedFetched = `-- name: MarkFeedFetched :one
UPDATE feeds
SET
last_fetched_at = $1,
updated_at = $1
WHERE id = $2
RETURNING id, name, url, created_at, updated_at, user_id, last_fetched_at
`

type MarkFeedFetchedParams struct {
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
	ID            uuid.UUID    `json:"id"`
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedFetched, arg.LastFetchedAt, arg.ID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}
