-- +goose Up
CREATE TABLE orders(
    id UUID primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    phone_num TEXT NOT NULL,
    email TEXT NOT NULL,
    look_up_code TEXT NOT NULL UNIQUE,
    number_of_wells INTEGER NOT NULL,
    predicted_full_price INTEGER NOT NULL
);

-- +goose Down
DROP TABLE orders;