-- name: CreateWell :one
INSERT INTO wells(id, created_at, updated_at, price_id, place_id, user_id, number)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: DeleteWells :exec
DELETE FROM wells;

-- name: DeleteWell :exec
DELETE FROM wells
WHERE id = $1;

-- name: DeleteWellByUser :exec
DELETE FROM wells
WHERE user_id = $1;

-- name: GetWells_ASC :many
SELECT * FROM wells
ORDER BY updated_at ASC;

-- name: GetWellsByUserID_ASC :many
SELECT * FROM wells
WHERE user_id = $1
ORDER BY created_at ASC;

-- name: GetWell :one
SELECT * FROM wells
WHERE id = $1;