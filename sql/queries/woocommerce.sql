-- name: CreateProduct :one
INSERT INTO woocommerce(
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
SELECT woocommerce.* FROM woocommerce
INNER JOIN users
ON woocommerce.user_id = users.id
WHERE user_id = $1
ORDER BY woocommerce.updated_at DESC
LIMIT $2 OFFSET $3;

-- name: SearchProductByTitle :many
SELECT * FROM woocommerce
WHERE "title" SIMILAR TO $1;

-- name: SearchProductBySKU :many
SELECT * FROM woocommerce
WHERE "sku" SIMILAR TO $1;