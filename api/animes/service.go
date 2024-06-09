package animes

import (
	"anime-list/database"
	"anime-list/database/queries"
	"context"
	"database/sql"
	"fmt"
	"strings"
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

func GetAnimesService(ctx context.Context, params GetAnimesParams) ([]Anime, error) {

	var animes []Anime
	var results []queries.Anime
	var err error

	if params.Query == "" {

		queryParams := queries.GetAnimesParams{
			Limit:  params.Limit,
			Offset: params.Offset,
		}

		results, err = database.Query.GetAnimes(ctx, queryParams)
		if err != nil {
			return animes, err
		}

	} else if params.Query != "" && !params.IncludeDescription {

		query := "%" + strings.TrimSpace(params.Query) + "%"

		queryParams := queries.SearchAnimesByTitleParams{
			TitleRomaji:  query,
			TitleNative:  sql.NullString{String: query, Valid: true},
			TitleEnglish: sql.NullString{String: query, Valid: true},
			Limit:        params.Limit,
			Offset:       params.Offset,
		}

		results, err = database.Query.SearchAnimesByTitle(ctx, queryParams)
		if err != nil {
			return animes, err
		}

	} else if params.Query != "" && params.IncludeDescription {

		query := "%" + strings.TrimSpace(params.Query) + "%"

		queryParams := queries.SearchAnimesByTitleAndDescriptionParams{
			TitleRomaji:  query,
			TitleNative:  sql.NullString{String: query, Valid: true},
			TitleEnglish: sql.NullString{String: query, Valid: true},
			Description:  query,
			Limit:        params.Limit,
			Offset:       params.Offset,
		}

		results, err = database.Query.SearchAnimesByTitleAndDescription(ctx, queryParams)
		if err != nil {
			return animes, err
		}

	}

	for _, anime := range results {
		element := Anime{}
		element.New(anime, []string{}, []string{}, CoverImage{})
		animes = append(animes, element)
	}

	return animes, nil
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

	coverImage, err := database.Query.GetCoverImageByAnimeId(ctx, result.ID)
	if err != nil {
		return anime, err
	}

	anime.New(result, genres, studios, NewCoverImage(&coverImage))

	return anime, nil
}

func UpdateAnime(id int) {}

func DeleteAnime(id int) {}
