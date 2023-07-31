-- name: CreateBookmark :one
INSERT INTO bookmarks (
    id,
    post_id,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBookmarkByPostID :one
SELECT * FROM bookmarks
WHERE post_id = $1;

-- name: SearchBookmarkByPostTitle :many
SELECT bookmarks.* FROM bookmarks
INNER JOIN posts
ON bookmarks.post_id = posts.id
WHERE posts.title SIMILAR TO $1;

-- name: GetBookmarkPostByUserDesc :many
SELECT * FROM bookmarks
WHERE user_id = $1
ORDER BY "updated_at" DESC
LIMIT $2 OFFSET $3;

-- name: GetBookmarkPostByUserAsc :many
SELECT * FROM bookmarks
WHERE user_id = $1
ORDER BY "updated_at" ASC
LIMIT $2 OFFSET $3;

-- name: RemoveBookmarkByPostID :one
DELETE FROM bookmarks
WHERE post_id = $1
RETURNING *;

-- name: RemoveBookmarkByUserID :one
DELETE FROM bookmarks
WHERE user_id = $1
RETURNING *;