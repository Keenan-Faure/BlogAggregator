-- name: CreateFeedFollow :one
INSERT INTO feed_follows(
    id,
    feed_id,
    user_id,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollow :one
SELECT * FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows;

-- name: GetUserFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = $1;

-- name: DeleteFeedByID :one
DELETE FROM feed_follows
WHERE id = $1 and user_id = $2
RETURNING *;
