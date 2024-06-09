-- name: CreateGenre :one
INSERT INTO genres (name)
VALUES (?)
RETURNING *;


-- name: GetGenre :one
SELECT * FROM genres
WHERE id = ?;


-- name: GetGenres :many
SELECT name FROM genres
WHERE id IN (
    SELECT genre_id FROM anime_genres
    WHERE anime_id = ?
);


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
