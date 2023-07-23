-- +goose Up
CREATE TABLE woocommerce(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    sku TEXT UNIQUE NOT NULL,
    price DECIMAL NOT NULL,
    qty INTEGER,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID,
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE woocommerce;