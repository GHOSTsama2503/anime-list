// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: studios.sql

package queries

import (
	"context"
)

const createStudio = `-- name: CreateStudio :one
INSERT INTO studios (name)
VALUES (?)
RETURNING id, name
`

func (q *Queries) CreateStudio(ctx context.Context, name string) (Studio, error) {
	row := q.db.QueryRowContext(ctx, createStudio, name)
	var i Studio
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const findStudio = `-- name: FindStudio :one
SELECT id, name FROM studios
WHERE name = ?
`

func (q *Queries) FindStudio(ctx context.Context, name string) (Studio, error) {
	row := q.db.QueryRowContext(ctx, findStudio, name)
	var i Studio
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}