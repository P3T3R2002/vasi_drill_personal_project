-- +goose Up
CREATE TABLE places(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    num_long INTEGER NOT NULL,
    num_vert INTEGER NOT NULL,
    expected_depth INTEGER NOT NULL
);

-- +goose Down
DROP TABLE places;