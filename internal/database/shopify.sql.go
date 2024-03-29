// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: shopify.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createShopifyProduct = `-- name: CreateShopifyProduct :one
INSERT INTO shopify(
    id,
    store_name,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, title, sku, price, qty, created_at, updated_at, store_name
`

type CreateShopifyProductParams struct {
	ID        uuid.UUID `json:"id"`
	StoreName string    `json:"store_name"`
	Title     string    `json:"title"`
	Sku       string    `json:"sku"`
	Price     string    `json:"price"`
	Qty       int32     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateShopifyProduct(ctx context.Context, arg CreateShopifyProductParams) (Shopify, error) {
	row := q.db.QueryRowContext(ctx, createShopifyProduct,
		arg.ID,
		arg.StoreName,
		arg.Title,
		arg.Sku,
		arg.Price,
		arg.Qty,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Shopify
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Sku,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.StoreName,
	)
	return i, err
}

const deleteTestShopifyProducts = `-- name: DeleteTestShopifyProducts :exec
DELETE FROM shopify
WHERE store_name = $1
`

func (q *Queries) DeleteTestShopifyProducts(ctx context.Context, storeName string) error {
	_, err := q.db.ExecContext(ctx, deleteTestShopifyProducts, storeName)
	return err
}

const getFirstRecordShopify = `-- name: GetFirstRecordShopify :one
SELECT id, title, sku, price, qty, created_at, updated_at, store_name FROM shopify
LIMIT 1
`

func (q *Queries) GetFirstRecordShopify(ctx context.Context) (Shopify, error) {
	row := q.db.QueryRowContext(ctx, getFirstRecordShopify)
	var i Shopify
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Sku,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.StoreName,
	)
	return i, err
}

const getShopifyProducts = `-- name: GetShopifyProducts :many
SELECT id, title, sku, price, qty, created_at, updated_at, store_name FROM shopify
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2
`

type GetShopifyProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetShopifyProducts(ctx context.Context, arg GetShopifyProductsParams) ([]Shopify, error) {
	rows, err := q.db.QueryContext(ctx, getShopifyProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shopify
	for rows.Next() {
		var i Shopify
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.StoreName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchShopifyProductByTitle = `-- name: SearchShopifyProductByTitle :many
SELECT id, title, sku, price, qty, created_at, updated_at, store_name FROM shopify
WHERE "title" SIMILAR TO $1
`

func (q *Queries) SearchShopifyProductByTitle(ctx context.Context, similarToEscape string) ([]Shopify, error) {
	rows, err := q.db.QueryContext(ctx, searchShopifyProductByTitle, similarToEscape)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shopify
	for rows.Next() {
		var i Shopify
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.StoreName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchShopifyShopifyProductBySKU = `-- name: SearchShopifyShopifyProductBySKU :many
SELECT id, title, sku, price, qty, created_at, updated_at, store_name FROM shopify
WHERE "sku" SIMILAR TO $1
`

func (q *Queries) SearchShopifyShopifyProductBySKU(ctx context.Context, similarToEscape string) ([]Shopify, error) {
	rows, err := q.db.QueryContext(ctx, searchShopifyShopifyProductBySKU, similarToEscape)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shopify
	for rows.Next() {
		var i Shopify
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.StoreName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateShopifyProducts = `-- name: UpdateShopifyProducts :one
UPDATE shopify SET
title = $1,
price = $2,
qty = $3,
updated_at = $4
WHERE "sku" = $5
RETURNING id, title, sku, price, qty, created_at, updated_at, store_name
`

type UpdateShopifyProductsParams struct {
	Title     string    `json:"title"`
	Price     string    `json:"price"`
	Qty       int32     `json:"qty"`
	UpdatedAt time.Time `json:"updated_at"`
	Sku       string    `json:"sku"`
}

func (q *Queries) UpdateShopifyProducts(ctx context.Context, arg UpdateShopifyProductsParams) (Shopify, error) {
	row := q.db.QueryRowContext(ctx, updateShopifyProducts,
		arg.Title,
		arg.Price,
		arg.Qty,
		arg.UpdatedAt,
		arg.Sku,
	)
	var i Shopify
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Sku,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.StoreName,
	)
	return i, err
}
