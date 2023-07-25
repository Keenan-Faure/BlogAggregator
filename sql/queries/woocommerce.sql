-- name: CreateWooProduct :one
INSERT INTO woocommerce(
    id,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: SearchWooProductByTitle :many
SELECT * FROM woocommerce
WHERE "title" SIMILAR TO $1;

-- name: SearchWooProductBySKU :many
SELECT * FROM woocommerce
WHERE "sku" SIMILAR TO $1;