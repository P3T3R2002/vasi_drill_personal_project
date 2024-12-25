-- name: CreateCode :one
INSERT INTO codes(code, created_at, updated_at, order_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;

-- name: GetCode :one
SELECT * FROM codes
WHERE code = $1;

-- name: DeleteCodes :exec
DELETE FROM codes;

-- name: DeleteCode :exec
DELETE FROM codes
WHERE code = $1;
