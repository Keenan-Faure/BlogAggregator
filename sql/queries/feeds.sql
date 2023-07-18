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

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE "url" = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at DESC;

-- name: MarkFeedFetched :one
UPDATE feeds
SET
last_fetched_at = $1,
updated_at = $1
WHERE id = $2
RETURNING *;
