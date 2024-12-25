-- +goose Up
CREATE TABLE grid(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    num_long FLOAT NOT NULL,
    num_vert FLOAT NOT NULL,
    expected_depth INTEGER NOT NULL
);

-- +goose Down
DROP TABLE grid;