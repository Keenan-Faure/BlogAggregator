-- name: CreateWooProduct :one
INSERT INTO woocommerce(
    id,
    store_name,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateWooProducts :one
UPDATE woocommerce SET
title = $1,
price = $2,
qty = $3,
updated_at = $4
WHERE "sku" = $5
RETURNING *;

-- name: SearchWooProductByTitle :many
SELECT * FROM woocommerce
WHERE "title" SIMILAR TO $1;

-- name: SearchWooProductBySKU :many
SELECT * FROM woocommerce
WHERE "sku" SIMILAR TO $1;

-- name: GetWooProducts :many
SELECT * FROM woocommerce
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;

-- name: GetFirstRecordWoo :one
SELECT * FROM woocommerce
LIMIT 1;

-- name: DeleteTestWooProducts :exec
DELETE FROM woocommerce
WHERE store_name = $1;