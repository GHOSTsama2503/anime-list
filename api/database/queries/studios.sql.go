// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
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

const getStudios = `-- name: GetStudios :many
SELECT name FROM studios
WHERE id IN (
    SELECT studio_id FROM anime_studios
    WHERE anime_id = ?
)
`

func (q *Queries) GetStudios(ctx context.Context, animeID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getStudios, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}