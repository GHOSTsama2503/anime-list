package animes

import (
	"anime-list/internal/database"
	"anime-list/internal/database/queries"
	"context"
	"database/sql"
	"fmt"
)

func CreateAnimeService(params CreateAnimeParams) (anime Anime, err error) {

	animeParams := queries.CreateAnimeParams{
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

	var transaction *sql.Tx
	if transaction, err = database.Db.Begin(); err != nil {
		return
	}

	query := database.Query.WithTx(transaction)
	context := context.Background()

	var animeDb queries.Anime
	if animeDb, err = query.CreateAnime(context, animeParams); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			return
		}
		return
	}

	var genres []string

	for _, name := range params.Genres {
		var result queries.Genre

		if result, err = query.CreateGenre(context, name); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
				return
			}
			return
		}

		animeGenreParams := queries.CreateAnimeGenreParams{
			AnimeID: animeDb.ID,
			GenreID: result.ID,
		}

		if _, err = query.CreateAnimeGenre(context, animeGenreParams); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
				return
			}
			return
		}

		genres = append(genres, result.Name)
	}

	var studios []string

	for _, name := range params.Studios {
		var result queries.Studio

		if result, err = query.CreateStudio(context, name); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
				return
			}
			return
		}

		animeStudioParams := queries.CreateAnimeStudioParams{
			AnimeID:  animeDb.ID,
			StudioID: result.ID,
		}

		if _, err = query.CreateAnimeStudio(context, animeStudioParams); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
				return
			}
			return
		}

		studios = append(studios, result.Name)
	}

	if err = transaction.Commit(); err != nil {
		return
	}

	anime.Id = int(animeDb.ID)
	anime.IdAl = int(animeDb.IDAl)
	anime.Title = Title{
		Romaji:  animeDb.TitleRomaji,
		Native:  animeDb.TitleNative.String,
		English: animeDb.TitleEnglish.String,
	}
	anime.Format = animeDb.Format
	anime.Status = animeDb.Status
	anime.Description = animeDb.Description
	anime.StartDate = anime.StartDate
	anime.EndDate = anime.EndDate
	anime.Season = animeDb.Season
	anime.SeasonYear = int(animeDb.SeasonYear.Int64)
	anime.Episodes = int(animeDb.Episodes)
	anime.Duration = int(animeDb.Duration)
	anime.Genres = genres
	anime.Studios = studios
	anime.CoverImage = params.CoverImage
	anime.BannerImage = animeDb.BannerImage.String

	return
}

func GetAnimesService(limit int64, offset int64) (animes []Anime, err error) {

	getAnimesParams := queries.GetAnimesParams{
		Limit:  limit,
		Offset: offset,
	}

	var results []queries.Anime
	if results, err = database.Query.GetAnimes(context.Background(), getAnimesParams); err != nil {
		return
	}

	for _, anime := range results {
		element := Anime{
			Id:   int(anime.ID),
			IdAl: int(anime.IDAl),
			Title: Title{
				Romaji:  anime.TitleRomaji,
				Native:  anime.TitleNative.String,
				English: anime.TitleEnglish.String,
			},
			Format:      anime.Format,
			Status:      anime.Status,
			Description: anime.Description,
			StartDate:   anime.StartDate,
			EndDate:     anime.EndDate,
			Season:      anime.Season,
			SeasonYear:  int(anime.SeasonYear.Int64),
			Episodes:    int(anime.Episodes),
			Duration:    int(anime.Duration),
			BannerImage: anime.BannerImage.String,
			StImage:     anime.StImage,
		}

		animes = append(animes, element)
	}

	return
}

func GetAnimeInfoService(id int) (anime Anime, err error) {
	return
}

func UpdateAnime(id int) {}

func DeleteAnime(id int) {}
