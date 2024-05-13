-- name: CreateAnimeStudio :one
INSERT INTO anime_studios (anime_id, studio_id)
VALUES ($1, $2)
RETURNING *;
