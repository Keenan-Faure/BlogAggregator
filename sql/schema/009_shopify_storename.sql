-- +goose Up
ALTER TABLE shopify ADD COLUMN store_name TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE shopify DROP COLUMN store_name;
