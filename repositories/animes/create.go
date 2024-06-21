package animes

import (
	"github.com/ghostsama2503/anime-list/api/repositories/animes/models"
	"context"
	"database/sql"
)

const create = `
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
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING
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
`

type CreateParams struct {
	IdAl         int64
	TitleRomaji  string
	TitleNative  sql.NullString
	TitleEnglish sql.NullString
	Format       string
	Status       string
	Description  string
	StartDate    string
	EndDate      string
	Season       string
	SeasonYear   sql.NullInt64
	Episodes     int64
	Duration     int64
	BannerImage  sql.NullString
	StImage      string
}

func (r *AnimesRepository) Create(ctx context.Context, params CreateParams) (models.Anime, error) {

	row := r.db.QueryRowContext(ctx, create,
		params.IdAl,
		params.TitleRomaji,
		params.TitleNative,
		params.TitleEnglish,
		params.Format,
		params.Status,
		params.Description,
		params.StartDate,
		params.EndDate,
		params.Season,
		params.SeasonYear,
		params.Episodes,
		params.Duration,
		params.BannerImage,
		params.StImage,
	)

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
