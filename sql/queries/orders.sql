-- name: CreateOrder :one
INSERT INTO orders(id, created_at, updated_at, name, phone_num, email, look_up_code, number_of_wells, predicted_full_price)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetOrderDetails :one
SELECT *  FROM orders
WHERE look_up_code = $1; 

-- name: GetCodes :many
SELECT look_up_code  FROM orders;

-- name: DeleteOrders :exec
DELETE FROM orders;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE look_up_code = $1;