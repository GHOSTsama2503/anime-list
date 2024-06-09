// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: groups.sql

package queries

import (
	"context"
)

const createAnimeGroup = `-- name: CreateAnimeGroup :one
INSERT INTO anime_groups (anime_id, group_id)
VALUES (?, ?)
RETURNING anime_id, group_id
`

type CreateAnimeGroupParams struct {
	AnimeID int64
	GroupID int64
}

func (q *Queries) CreateAnimeGroup(ctx context.Context, arg CreateAnimeGroupParams) (AnimeGroup, error) {
	row := q.db.QueryRowContext(ctx, createAnimeGroup, arg.AnimeID, arg.GroupID)
	var i AnimeGroup
	err := row.Scan(&i.AnimeID, &i.GroupID)
	return i, err
}

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES (?)
RETURNING id, name
`

func (q *Queries) CreateGroup(ctx context.Context, name string) (Group, error) {
	row := q.db.QueryRowContext(ctx, createGroup, name)
	var i Group
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const findGroup = `-- name: FindGroup :one
SELECT id, name FROM groups
WHERE name = ?
`

func (q *Queries) FindGroup(ctx context.Context, name string) (Group, error) {
	row := q.db.QueryRowContext(ctx, findGroup, name)
	var i Group
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getGroup = `-- name: GetGroup :one
SELECT id, name FROM groups
WHERE id = ?
`

func (q *Queries) GetGroup(ctx context.Context, id int64) (Group, error) {
	row := q.db.QueryRowContext(ctx, getGroup, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
