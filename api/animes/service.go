package animes

import (
	"anime-list/database"
	"anime-list/database/queries"
	"context"
	"database/sql"
	"fmt"
)

func CreateAnimeService(params CreateAnimeParams) (anime Anime, err error) {

	var transaction *sql.Tx
	if transaction, err = database.Db.Begin(); err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		}
	}()

	query := database.Query.WithTx(transaction)
	context := context.Background()

	animeParams := newAnimeParams(params)

	var group queries.Group

	if params.Group != "" {
		group, _ = query.FindGroup(context, params.Group)

		if group.ID == 0 {
			if group, err = query.CreateGroup(context, params.Group); err != nil {
				return
			}
		}

		animeParams.GroupPosition = sql.NullInt64{Int64: int64(params.GroupPosition), Valid: true}
	}

	var animeDb queries.Anime
	if animeDb, err = query.CreateAnime(context, animeParams); err != nil {
		return
	}

	if group.ID != 0 {
		animeGroupParams := queries.CreateAnimeGroupParams{
			AnimeID: animeDb.ID,
			GroupID: group.ID,
		}

		if _, err = query.CreateAnimeGroup(context, animeGroupParams); err != nil {
			return
		}
	}

	var genres []string
	if genres, err = insertGenres(animeDb.ID, params.Genres, query, context); err != nil {
		return
	}

	var studios []string
	if studios, err = insertStudios(animeDb.ID, params.Studios, query, context); err != nil {
		return
	}

	if err = transaction.Commit(); err != nil {
		return
	}

	anime = newAnimeFromParams(animeDb, genres, studios, params)

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

func GetAnimeInfoService(ctx context.Context, id int) (Anime, error) {
	var anime Anime

	result, err := database.Query.GetAnime(ctx, int64(id))
	if err != nil {
		return anime, err
	}

	genres, err := database.Query.GetGenres(ctx, result.ID)
	if err != nil {
		return anime, err
	}

	studios, err := database.Query.GetStudios(ctx, result.ID)
	if err != nil {
		return anime, err
	}

	anime.New(result, genres, studios, CoverImage{})

	return anime, nil
}

func UpdateAnime(id int) {}

func DeleteAnime(id int) {}
