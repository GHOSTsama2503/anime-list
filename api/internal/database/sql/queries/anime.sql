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
    st_image,
    group_position
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
    ?,
    ?
)
RETURNING *;

-- name: GetAnimes :many
SELECT * FROM animes
ORDER BY title_romaji COLLATE NOCASE ASC
LIMIT ?
OFFSET ?;


-- name: GetAnime :one
SELECT * FROM animes
WHERE id = ?;
