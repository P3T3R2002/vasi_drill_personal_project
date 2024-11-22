-- +goose Up
CREATE TABLE wells(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    place_id UUID NOT NULL REFERENCES places(id),
    price_id UUID NOT NULL REFERENCES prices(id),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    number INTEGER NOT NULL
);

-- +goose Down
DROP TABLE wells;