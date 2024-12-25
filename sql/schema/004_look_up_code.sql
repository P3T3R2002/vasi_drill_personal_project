-- +goose Up
CREATE TABLE codes(
    code TEXT primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE codes;