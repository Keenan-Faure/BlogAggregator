-- name: CreateLike :one
INSERT INTO liked (
    id,
    post_id,
    user_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetLikedByPostID :one
SELECT * FROM liked
WHERE post_id = $1;

-- name: SearchLikedByPostTitle :many
SELECT liked.* FROM liked
INNER JOIN posts
ON liked.post_id = posts.id
WHERE posts.title = $1;

-- name: RemoveLikedByPostID :one
DELETE FROM liked
WHERE post_id = $1
RETURNING *;

-- name: RemoveLikedByUserID :one
DELETE FROM liked
WHERE user_id = $1
RETURNING *;