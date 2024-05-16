package animes

import (
	"anime-list/internal/database/queries"
	"context"
)

func newAnimeParams(params CreateAnimeParams) (dv queries.CreateAnimeParams) {
	return queries.CreateAnimeParams{
		IDAl:         int64(params.IdAl),
		TitleRomaji:  params.Title.Romaji,
		TitleNative:  NullString(params.Title.Native),
		TitleEnglish: NullString(params.Title.English),
		Format:       params.Format,
		Status:       params.Status,
		Description:  params.Description,
		StartDate:    FuzzyDateToString(params.StartDate),
		EndDate:      FuzzyDateToString(params.EndDate),
		Season:       string(params.Season),
		SeasonYear:   NullInt64(int64(params.SeasonYear)),
		Episodes:     int64(params.Episodes),
		Duration:     int64(params.Duration),
		BannerImage:  NullString(params.BannerImage),
		StImage:      StImage(params.IdAl),
	}
}

func newAnimeFromParams(animeDb queries.Anime, genres []string, studios []string, params CreateAnimeParams) Anime {
	return Anime{
		Id:   int(animeDb.ID),
		IdAl: int(animeDb.IDAl),
		Title: Title{
			Romaji:  animeDb.TitleRomaji,
			Native:  animeDb.TitleNative.String,
			English: animeDb.TitleEnglish.String,
		},
		Format:      animeDb.Format,
		Status:      animeDb.Status,
		Description: animeDb.Description,
		StartDate:   animeDb.StartDate,
		EndDate:     animeDb.EndDate,
		Season:      animeDb.Season,
		SeasonYear:  int(animeDb.SeasonYear.Int64),
		Episodes:    int(animeDb.Episodes),
		Duration:    int(animeDb.Duration),
		Genres:      genres,
		Studios:     studios,
		CoverImage:  params.CoverImage,
		BannerImage: animeDb.BannerImage.String,
		StImage:     StImage(int(animeDb.ID)),
	}
}

func insertGenres(animeId int64, genres []string, query *queries.Queries, ctx context.Context) (ret []string, err error) {

	for _, name := range genres {
		var result queries.Genre

		result, _ = query.FindGenre(ctx, name)

		if result.ID == 0 {
			if result, err = query.CreateGenre(ctx, name); err != nil {
				return
			}
		}

		animeGenreParams := queries.CreateAnimeGenreParams{
			AnimeID: animeId,
			GenreID: result.ID,
		}

		if _, err = query.CreateAnimeGenre(ctx, animeGenreParams); err != nil {
			return
		}

		ret = append(ret, result.Name)
	}

	return
}

func insertStudios(animeId int64, studios []string, query *queries.Queries, ctx context.Context) (ret []string, err error) {

	for _, name := range studios {
		var result queries.Studio

		result, _ = query.FindStudio(ctx, name)

		if result.ID == 0 {
			if result, err = query.CreateStudio(ctx, name); err != nil {
				return
			}
		}

		animeStudioParams := queries.CreateAnimeStudioParams{
			AnimeID:  animeId,
			StudioID: result.ID,
		}

		if _, err = query.CreateAnimeStudio(ctx, animeStudioParams); err != nil {
			return
		}

		ret = append(ret, result.Name)
	}

	return
}
