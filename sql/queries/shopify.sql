-- name: CreateProduct :one
INSERT INTO shopify(
    id,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at.
    user_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetProductByUser :many
SELECT shopify.* FROM shopify
INNER JOIN users
ON shopify.user_id = users.id
WHERE user_id = $1
ORDER BY shopify.updated_at DESC
LIMIT $2 OFFSET $3;

-- name: SearchProductByTitle :many
SELECT * FROM shopify
WHERE "title" SIMILAR TO $1;

-- name: SearchProductBySKU :many
SELECT * FROM shopify
WHERE "sku" SIMILAR TO $1;

