-- +goose Up
CREATE TABLE wells(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    gps_long FLOAT NOT NULL,
    gps_vert FLOAT NOT NULL,
    cell_id UUID NOT NULL REFERENCES grid(id) ON DELETE CASCADE,
    price INTEGER NOT NULL,
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE wells;