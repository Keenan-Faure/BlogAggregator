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

-- name: SearchShopifyProductByTitle :many
SELECT * FROM shopify
WHERE "title" SIMILAR TO $1;

-- name: SearchShopifyShopifyProductBySKU :many
SELECT * FROM shopify
WHERE "sku" SIMILAR TO $1;

