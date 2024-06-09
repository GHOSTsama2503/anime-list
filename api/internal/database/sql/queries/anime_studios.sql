-- name: CreateAnimeStudio :one
INSERT INTO anime_studios (anime_id, studio_id)
VALUES (?, ?)
RETURNING *;
