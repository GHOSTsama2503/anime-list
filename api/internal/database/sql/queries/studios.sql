-- name: CreateStudio :one
INSERT INTO studios (name)
VALUES (?)
RETURNING *;
