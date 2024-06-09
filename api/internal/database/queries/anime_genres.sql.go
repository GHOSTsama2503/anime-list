// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_genres.sql

package queries

import (
	"context"
)

const createAnimeGenre = `-- name: CreateAnimeGenre :one
INSERT INTO anime_genres (anime_id, genre_id)
VALUES (?, ?)
RETURNING anime_id, genre_id
`

type CreateAnimeGenreParams struct {
	AnimeID int64
	GenreID int64
}

func (q *Queries) CreateAnimeGenre(ctx context.Context, arg CreateAnimeGenreParams) (AnimeGenre, error) {
	row := q.db.QueryRowContext(ctx, createAnimeGenre, arg.AnimeID, arg.GenreID)
	var i AnimeGenre
	err := row.Scan(&i.AnimeID, &i.GenreID)
	return i, err
}
