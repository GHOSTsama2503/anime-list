-- name: CreateStudio :one
INSERT INTO studios (name)
VALUES (?)
RETURNING *;


-- name: GetStudios :many
SELECT name FROM studios
WHERE id IN (
    SELECT studio_id FROM anime_studios
    WHERE anime_id = ?
);


-- name: FindStudio :one
SELECT * FROM studios
WHERE name = ?;
