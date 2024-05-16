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
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetAnimes :many
SELECT * FROM animes
LIMIT ?
OFFSET ?;
