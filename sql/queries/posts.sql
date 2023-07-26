-- name: CreatePost :one
INSERT INTO posts(
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET
"updated_at" = $1,
"title" = $2,
"description" = $3
WHERE "url" = $4
RETURNING *;

-- name: GetPostsByUserAsc :many
SELECT posts.*
FROM posts
JOIN feed_follows
ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at ASC
LIMIT $2 OFFSET $3;

-- name: GetPostsByUserDesc :many
SELECT posts.*
FROM posts
JOIN feed_follows
ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2 OFFSET $3;

-- name: GetPostSearchTitle :many
SELECT * FROM posts WHERE "title"
SIMILAR TO $1;

-- name: GetFirstRecordPost :one
SELECT * FROM posts
LIMIT 1;

-- name: DeleteTestPosts :exec
DELETE FROM posts
WHERE feed_id = $1;