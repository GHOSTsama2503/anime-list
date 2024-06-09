-- name: CreateGenre :one
INSERT INTO genres (name)
VALUES (?)
RETURNING *;


-- name: GetGenre :one
SELECT * FROM genres
WHERE id = ?;


-- name: FindGenre :one
SELECT * FROM genres
WHERE name = ?;


-- name: UpdateGenre :one
UPDATE genres
SET name = ?
WHERE id = ?
RETURNING *;


-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = ?;
