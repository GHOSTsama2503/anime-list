-- name: CreateCoverImage :one
INSERT INTO cover_images (anime_id, extra_large, large, medium, color)
VALUES (?, ?, ?, ?, ?)
RETURNING *;


-- name: GetCoverImageByAnimeId :one
SELECT * FROM cover_images
WHERE anime_id = ?;


-- name: DeleteCoverImage :exec
DELETE FROM cover_images
WHERE id = ?;


-- name: DeleteCoverImageByAnimeId :exec
DELETE FROM cover_images
WHERE anime_id = ?;
