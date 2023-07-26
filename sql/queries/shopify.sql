-- name: CreateShopifyProduct :one
INSERT INTO shopify(
    id,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateShopifyProducts :one
UPDATE shopify SET
title = $1,
price = $2,
qty = $3,
updated_at = $4
WHERE "sku" = $5
RETURNING *;

-- name: SearchShopifyProductByTitle :many
SELECT * FROM shopify
WHERE "title" SIMILAR TO $1;

-- name: SearchShopifyShopifyProductBySKU :many
SELECT * FROM shopify
WHERE "sku" SIMILAR TO $1;

-- name: GetShopifyProducts :many
SELECT * FROM shopify
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;

-- >> used for tests << --

