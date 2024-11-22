-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, email, hashed_password)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdatePassword :one
UPDATE users
SET updated_at = NOW(), email = $2, hashed_password = $3
WHERE id = $1
RETURNING *;
