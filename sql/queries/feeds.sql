-- name: CreateFeed :one
INSERT INTO feeds(
    id,
    name,
    url,
    created_at,
    updated_at,
    last_fetched_at,
    user_id
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFeedsAsc :many
SELECT * FROM feeds
ORDER BY updated_at ASC
LIMIT $1 OFFSET $2;

-- name: GetFeedsDesc :many
SELECT * FROM feeds
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE "url" = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at DESC
LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET
last_fetched_at = $1,
updated_at = $1
WHERE id = $2
RETURNING *;

-- name: GetFeedSearchName :many
SELECT * FROM feeds WHERE "name"
SIMILAR TO $1;