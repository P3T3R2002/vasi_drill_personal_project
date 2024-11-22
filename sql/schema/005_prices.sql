-- +goose Up
CREATE TABLE prices(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    radious INTEGER NOT NULL,
    price_per_meter INTEGER NOT NULL
);

-- +goose Down
DROP TABLE prices;