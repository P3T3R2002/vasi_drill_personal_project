// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: grid.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createCell = `-- name: CreateCell :one
INSERT INTO grid(id, created_at, updated_at, num_long, num_vert, expected_depth)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, num_long, num_vert, expected_depth
`

type CreateCellParams struct {
	ID            uuid.UUID
	NumLong       float64
	NumVert       float64
	ExpectedDepth int32
}

func (q *Queries) CreateCell(ctx context.Context, arg CreateCellParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, createCell,
		arg.ID,
		arg.NumLong,
		arg.NumVert,
		arg.ExpectedDepth,
	)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}

const deleteGrid = `-- name: DeleteGrid :exec
DELETE FROM grid
`

func (q *Queries) DeleteGrid(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteGrid)
	return err
}

const getClosestGrid_BL = `-- name: GetClosestGrid_BL :one
SELECT id, created_at, updated_at, num_long, num_vert, expected_depth FROM grid
WHERE num_long >= $1 AND num_vert <= $2
ORDER BY num_long ASC, num_vert DESC
LIMIT 1
`

type GetClosestGrid_BLParams struct {
	NumLong float64
	NumVert float64
}

func (q *Queries) GetClosestGrid_BL(ctx context.Context, arg GetClosestGrid_BLParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, getClosestGrid_BL, arg.NumLong, arg.NumVert)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}

const getClosestGrid_BR = `-- name: GetClosestGrid_BR :one
SELECT id, created_at, updated_at, num_long, num_vert, expected_depth FROM grid
WHERE num_long >= $1 AND num_vert >= $2
ORDER BY num_long ASC, num_vert ASC
LIMIT 1
`

type GetClosestGrid_BRParams struct {
	NumLong float64
	NumVert float64
}

func (q *Queries) GetClosestGrid_BR(ctx context.Context, arg GetClosestGrid_BRParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, getClosestGrid_BR, arg.NumLong, arg.NumVert)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}

const getClosestGrid_TL = `-- name: GetClosestGrid_TL :one
SELECT id, created_at, updated_at, num_long, num_vert, expected_depth FROM grid
WHERE num_long <= $1 AND num_vert <= $2
ORDER BY num_long DESC, num_vert DESC
LIMIT 1
`

type GetClosestGrid_TLParams struct {
	NumLong float64
	NumVert float64
}

func (q *Queries) GetClosestGrid_TL(ctx context.Context, arg GetClosestGrid_TLParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, getClosestGrid_TL, arg.NumLong, arg.NumVert)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}

const getClosestGrid_TR = `-- name: GetClosestGrid_TR :one
SELECT id, created_at, updated_at, num_long, num_vert, expected_depth FROM grid
WHERE num_long <= $1 AND num_vert >= $2
ORDER BY num_long DESC, num_vert ASC
LIMIT 1
`

type GetClosestGrid_TRParams struct {
	NumLong float64
	NumVert float64
}

func (q *Queries) GetClosestGrid_TR(ctx context.Context, arg GetClosestGrid_TRParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, getClosestGrid_TR, arg.NumLong, arg.NumVert)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}

const getDepthByGridNum = `-- name: GetDepthByGridNum :one
SELECT grid.expected_depth FROM grid
WHERE num_long = $1 AND num_vert = $2
`

type GetDepthByGridNumParams struct {
	NumLong float64
	NumVert float64
}

func (q *Queries) GetDepthByGridNum(ctx context.Context, arg GetDepthByGridNumParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getDepthByGridNum, arg.NumLong, arg.NumVert)
	var expected_depth int32
	err := row.Scan(&expected_depth)
	return expected_depth, err
}

const updateCell = `-- name: UpdateCell :one
UPDATE grid
SET updated_at = NOW(), expected_depth = $3
WHERE num_long = $1 AND num_vert = $2
RETURNING id, created_at, updated_at, num_long, num_vert, expected_depth
`

type UpdateCellParams struct {
	NumLong       float64
	NumVert       float64
	ExpectedDepth int32
}

func (q *Queries) UpdateCell(ctx context.Context, arg UpdateCellParams) (Grid, error) {
	row := q.db.QueryRowContext(ctx, updateCell, arg.NumLong, arg.NumVert, arg.ExpectedDepth)
	var i Grid
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.NumLong,
		&i.NumVert,
		&i.ExpectedDepth,
	)
	return i, err
}
