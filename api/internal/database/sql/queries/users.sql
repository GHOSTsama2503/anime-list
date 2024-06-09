-- name: CreateUser :one
INSERT INTO users (username, password, is_admin)
VALUES (?, ?, ?)
RETURNING *;


-- name: CheckUser :many
SELECT * FROM users
WHERE username = ? AND password = ?;


-- name: GetUserByUsername :many
SELECT * FROM users
WHERE username = ?;


-- name: GetUserById :many
SELECT * FROM users
WHERE id = ?;


-- name: AdminExists :one
SELECT COUNT(is_admin) FROM users;
