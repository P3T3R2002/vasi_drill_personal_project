-- name: CreateCell :one
INSERT INTO grid(id, created_at, updated_at, num_long, num_vert, expected_depth)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteGrid :exec
DELETE FROM grid;

-- name: GetDepthByGridNum :one
SELECT grid.expected_depth FROM grid
WHERE num_long = $1 AND num_vert = $2;

-- name: GetClosestGrid_TL :one
SELECT * FROM grid
WHERE num_long <= $1 AND num_vert <= $2
ORDER BY num_long DESC, num_vert DESC
LIMIT 1;

-- name: GetClosestGrid_TR :one
SELECT * FROM grid
WHERE num_long <= $1 AND num_vert >= $2
ORDER BY num_long DESC, num_vert ASC
LIMIT 1;

-- name: GetClosestGrid_BL :one
SELECT * FROM grid
WHERE num_long >= $1 AND num_vert <= $2
ORDER BY num_long ASC, num_vert DESC
LIMIT 1;

-- name: GetClosestGrid_BR :one
SELECT * FROM grid
WHERE num_long >= $1 AND num_vert >= $2
ORDER BY num_long ASC, num_vert ASC
LIMIT 1;

-- name: UpdateCell :one
UPDATE grid
SET updated_at = NOW(), expected_depth = $3
WHERE num_long = $1 AND num_vert = $2
RETURNING *;
