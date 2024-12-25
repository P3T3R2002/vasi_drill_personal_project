// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
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
RETURNING id, created_at, updated_at, name, phone_num, email, look_up_code, number_of_wells, predicted_full_price
`

type CreateOrderParams struct {
	ID                 uuid.UUID
	Name               string
	PhoneNum           string
	Email              string
	LookUpCode         string
	NumberOfWells      int32
	PredictedFullPrice int32
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.Name,
		arg.PhoneNum,
		arg.Email,
		arg.LookUpCode,
		arg.NumberOfWells,
		arg.PredictedFullPrice,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.PhoneNum,
		&i.Email,
		&i.LookUpCode,
		&i.NumberOfWells,
		&i.PredictedFullPrice,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE look_up_code = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, lookUpCode string) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, lookUpCode)
	return err
}

const deleteOrders = `-- name: DeleteOrders :exec
DELETE FROM orders
`

func (q *Queries) DeleteOrders(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteOrders)
	return err
}

const getCodes = `-- name: GetCodes :many
SELECT look_up_code  FROM orders
`

func (q *Queries) GetCodes(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getCodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var look_up_code string
		if err := rows.Scan(&look_up_code); err != nil {
			return nil, err
		}
		items = append(items, look_up_code)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderDetails = `-- name: GetOrderDetails :one
SELECT id, created_at, updated_at, name, phone_num, email, look_up_code, number_of_wells, predicted_full_price  FROM orders
WHERE look_up_code = $1
`

func (q *Queries) GetOrderDetails(ctx context.Context, lookUpCode string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderDetails, lookUpCode)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.PhoneNum,
		&i.Email,
		&i.LookUpCode,
		&i.NumberOfWells,
		&i.PredictedFullPrice,
	)
	return i, err
}
