-- +goose Up
CREATE TABLE shopify(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    sku TEXT UNIQUE NOT NULL,
    price DECIMAL NOT NULL,
    qty INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE shopify;