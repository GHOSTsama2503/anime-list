-- name: CreateStudio :one
INSERT INTO studios (name)
VALUES (?)
RETURNING *;


-- name: FindStudio :one
SELECT * FROM studios
WHERE name = ?;
