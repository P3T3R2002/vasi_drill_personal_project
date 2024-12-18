// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: wells.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createWell = `-- name: CreateWell :one
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
RETURNING id, created_at, updated_at, place_id, price_id, user_id, number
`

type CreateWellParams struct {
	ID      uuid.UUID
	PriceID uuid.UUID
	PlaceID uuid.UUID
	UserID  uuid.UUID
	Number  int32
}

func (q *Queries) CreateWell(ctx context.Context, arg CreateWellParams) (Well, error) {
	row := q.db.QueryRowContext(ctx, createWell,
		arg.ID,
		arg.PriceID,
		arg.PlaceID,
		arg.UserID,
		arg.Number,
	)
	var i Well
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PlaceID,
		&i.PriceID,
		&i.UserID,
		&i.Number,
	)
	return i, err
}

const deleteWell = `-- name: DeleteWell :exec
DELETE FROM wells
WHERE id = $1
`

func (q *Queries) DeleteWell(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteWell, id)
	return err
}

const deleteWellByUser = `-- name: DeleteWellByUser :exec
DELETE FROM wells
WHERE user_id = $1
`

func (q *Queries) DeleteWellByUser(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteWellByUser, userID)
	return err
}

const deleteWells = `-- name: DeleteWells :exec
DELETE FROM wells
`

func (q *Queries) DeleteWells(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteWells)
	return err
}

const getWell = `-- name: GetWell :one
SELECT id, created_at, updated_at, place_id, price_id, user_id, number FROM wells
WHERE id = $1
`

func (q *Queries) GetWell(ctx context.Context, id uuid.UUID) (Well, error) {
	row := q.db.QueryRowContext(ctx, getWell, id)
	var i Well
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PlaceID,
		&i.PriceID,
		&i.UserID,
		&i.Number,
	)
	return i, err
}

const getWellsByUserID_ASC = `-- name: GetWellsByUserID_ASC :many
SELECT id, created_at, updated_at, place_id, price_id, user_id, number FROM wells
WHERE user_id = $1
ORDER BY created_at ASC
`

func (q *Queries) GetWellsByUserID_ASC(ctx context.Context, userID uuid.UUID) ([]Well, error) {
	rows, err := q.db.QueryContext(ctx, getWellsByUserID_ASC, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Well
	for rows.Next() {
		var i Well
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PlaceID,
			&i.PriceID,
			&i.UserID,
			&i.Number,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWells_ASC = `-- name: GetWells_ASC :many
SELECT id, created_at, updated_at, place_id, price_id, user_id, number FROM wells
ORDER BY updated_at ASC
`

func (q *Queries) GetWells_ASC(ctx context.Context) ([]Well, error) {
	rows, err := q.db.QueryContext(ctx, getWells_ASC)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Well
	for rows.Next() {
		var i Well
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PlaceID,
			&i.PriceID,
			&i.UserID,
			&i.Number,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
