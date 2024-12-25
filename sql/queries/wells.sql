-- name: CreateWell :one
INSERT INTO wells(id, created_at, updated_at, gps_long, gps_vert, price, cell_id, order_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5, 
    $6
)
RETURNING *;

-- name: GetWellDetails :many
SELECT w.id, w.gps_long, w.gps_vert, g.expected_depth, w.price FROM wells w
INNER JOIN grid g ON g.id = w.cell_id
WHERE w.order_id = $1;

-- name: DeleteWells :exec
DELETE FROM wells;

-- name: DeleteWell :exec
DELETE FROM wells
WHERE id = $1;
