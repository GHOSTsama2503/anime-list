-- name: CreateAnimeGenre :one
INSERT INTO anime_genres (anime_id, genre_id)
VALUES ($1, $2)
RETURNING *;
