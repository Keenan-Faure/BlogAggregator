// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: woocommerce.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createWooProduct = `-- name: CreateWooProduct :one
INSERT INTO woocommerce(
    id,
    title,
    sku,
    price,
    qty,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, sku, price, qty, created_at, updated_at
`

type CreateWooProductParams struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Sku       string    `json:"sku"`
	Price     string    `json:"price"`
	Qty       int32     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateWooProduct(ctx context.Context, arg CreateWooProductParams) (Woocommerce, error) {
	row := q.db.QueryRowContext(ctx, createWooProduct,
		arg.ID,
		arg.Title,
		arg.Sku,
		arg.Price,
		arg.Qty,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Woocommerce
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Sku,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getWooProducts = `-- name: GetWooProducts :many
SELECT id, title, sku, price, qty, created_at, updated_at FROM woocommerce
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2
`

type GetWooProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetWooProducts(ctx context.Context, arg GetWooProductsParams) ([]Woocommerce, error) {
	rows, err := q.db.QueryContext(ctx, getWooProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Woocommerce
	for rows.Next() {
		var i Woocommerce
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const searchWooProductBySKU = `-- name: SearchWooProductBySKU :many
SELECT id, title, sku, price, qty, created_at, updated_at FROM woocommerce
WHERE "sku" SIMILAR TO $1
`

func (q *Queries) SearchWooProductBySKU(ctx context.Context, similarToEscape string) ([]Woocommerce, error) {
	rows, err := q.db.QueryContext(ctx, searchWooProductBySKU, similarToEscape)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Woocommerce
	for rows.Next() {
		var i Woocommerce
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const searchWooProductByTitle = `-- name: SearchWooProductByTitle :many
SELECT id, title, sku, price, qty, created_at, updated_at FROM woocommerce
WHERE "title" SIMILAR TO $1
`

func (q *Queries) SearchWooProductByTitle(ctx context.Context, similarToEscape string) ([]Woocommerce, error) {
	rows, err := q.db.QueryContext(ctx, searchWooProductByTitle, similarToEscape)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Woocommerce
	for rows.Next() {
		var i Woocommerce
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Sku,
			&i.Price,
			&i.Qty,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateWooProducts = `-- name: UpdateWooProducts :one
UPDATE woocommerce SET
title = $1,
price = $2,
qty = $3,
updated_at = $4
WHERE "sku" = $5
RETURNING id, title, sku, price, qty, created_at, updated_at
`

type UpdateWooProductsParams struct {
	Title     string    `json:"title"`
	Price     string    `json:"price"`
	Qty       int32     `json:"qty"`
	UpdatedAt time.Time `json:"updated_at"`
	Sku       string    `json:"sku"`
}

func (q *Queries) UpdateWooProducts(ctx context.Context, arg UpdateWooProductsParams) (Woocommerce, error) {
	row := q.db.QueryRowContext(ctx, updateWooProducts,
		arg.Title,
		arg.Price,
		arg.Qty,
		arg.UpdatedAt,
		arg.Sku,
	)
	var i Woocommerce
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Sku,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
