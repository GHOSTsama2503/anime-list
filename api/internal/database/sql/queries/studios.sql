-- name: CreateStudio :one
INSERT INTO studios (name)
VALUES ($1)
RETURNING *;
