package animes

import (
	"anime-list/repositories/animes/models"
	"context"
)

const get = `
SELECT
	id,
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
FROM animes
WHERE id = ?;
`

func (q *AnimesRepository) Get(ctx context.Context, id int64) (models.Anime, error) {
	row := q.db.QueryRowContext(ctx, get, id)

	var anime models.Anime

	err := row.Scan(
		&anime.Id,
		&anime.IdAl,
		&anime.TitleRomaji,
		&anime.TitleNative,
		&anime.TitleEnglish,
		&anime.Format,
		&anime.Status,
		&anime.Description,
		&anime.StartDate,
		&anime.EndDate,
		&anime.Season,
		&anime.SeasonYear,
		&anime.Episodes,
		&anime.Duration,
		&anime.BannerImage,
		&anime.StImage,
	)

	return anime, err
}
