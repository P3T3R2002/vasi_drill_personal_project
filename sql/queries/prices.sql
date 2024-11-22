-- name: CreatePrice :one
INSERT INTO prices (id, created_at, updated_at, name, radious, price_per_meter)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeletePrices :exec
DELETE FROM users;

-- name: GetPriceByRadious :one
SELECT * FROM prices
WHERE radious = $1;

-- name: UpdatePrice :one
UPDATE prices
SET updated_at = NOW(), price_per_meter = $2
WHERE id = $1
RETURNING *;
