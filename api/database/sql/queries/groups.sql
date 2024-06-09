-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES (?)
RETURNING *;


-- name: FindGroup :one
SELECT * FROM groups
WHERE name = ?;


-- name: GetGroup :one
SELECT * FROM groups
WHERE id = ?;


-- name: CreateAnimeGroup :one
INSERT INTO anime_groups (anime_id, group_id)
VALUES (?, ?)
RETURNING *;
