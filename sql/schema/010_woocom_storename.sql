-- +goose Up
ALTER TABLE woocommerce ADD COLUMN store_name TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE woocommerce DROP COLUMN store_name;
