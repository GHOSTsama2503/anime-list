-- name: CreateAnime :one
INSERT INTO animes (
    id_al,
    title_romaji,
    title_native,
    title_english,
    format,
    status,
    description,
    start_date,
    end_date,
    season,
    season_year,
    episodes,
    duration,
    banner_image,
    st_image
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
)
RETURNING *;

-- name: GetAnimes :many
SELECT * FROM animes
LIMIT $1
OFFSET $2;
