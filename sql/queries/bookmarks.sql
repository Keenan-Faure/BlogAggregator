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
WHERE posts.title = $1;

-- name: RemoveBookmarkByPostID :one
DELETE FROM bookmarks
WHERE post_id = $1
RETURNING *;

-- name: RemoveBookmarkByUserID :one
DELETE FROM bookmarks
WHERE user_id = $1
RETURNING *;