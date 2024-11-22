-- +goose Up
CREATE TABLE users(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    hashed_password  TEXT NOT NULL DEFAULT 'unset'
);

-- +goose Down
DROP TABLE users;