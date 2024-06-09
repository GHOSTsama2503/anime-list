package animes

import (
	"anime-list/anilist"
	"anime-list/common/config"
	"context"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

func RemoteSearchController(context context.Context, input *RemoteSearchRequest) (*RemoteSearchResponse, error) {
	response := &RemoteSearchResponse{}

	results, err := anilist.SearchAnime(input.Body.Title, 0, 20)
	if err != nil {
		return response, err
	}

	responseResults := []RemoteSearchResult{}
	for _, value := range results {
		element := RemoteSearchResult{}
		element.Id = value.Id
		element.Image = value.Image.Medium
		element.Title = value.Title
		responseResults = append(responseResults, element)
	}

	response.Body = responseResults

	return response, nil
}

func CreateAnimeController(context context.Context, input *CreateAnimeRequest) (*CreateAnimeResponse, error) {
	response := &CreateAnimeResponse{}

	result, err := anilist.GetAnime(input.Body.Id)
	if err != nil {
		return response, err
	}

	var studios []string
	for _, s := range result.Studios.Nodes {
		studios = append(studios, s.Name)
	}

	params := CreateAnimeParams{
		IdAl:          int(result.Id),
		Title:         result.Title,
		Format:        result.Format,
		Status:        result.Status,
		Description:   result.Description,
		StartDate:     FuzzyDate(result.StartDate),
		EndDate:       FuzzyDate(result.EndDate),
		Season:        Season(result.Season),
		SeasonYear:    result.SeasonYear,
		Episodes:      result.Episodes,
		Duration:      result.Duration,
		CoverImage:    CoverImage(result.CoverImage),
		BannerImage:   result.BannerImage,
		Genres:        result.Genres,
		Studios:       studios,
		Group:         input.Body.Group,
		GroupPosition: input.Body.GroupPosition,
	}

	_, err = CreateAnimeService(params)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %v", uuid.New(), err.Error()))
		return response, err
	}

	response.Body.Ok = true

	return response, nil
}

func GetAnimesController(context context.Context, input *GetAnimesRequest) (*GetAnimesResponse, error) {
	response := &GetAnimesResponse{}

	if input.Limit <= 0 {
		input.Limit = config.Env.DefaultLimit
	}

	if input.Limit > config.Env.MaxLimit {
		log.Warnf("request limit (%d) is greater than the max allowed, using fallback: '%d'", input.Limit, config.Env.DefaultLimit)
		input.Limit = config.Env.DefaultLimit
	}

	if input.Offset < 0 {
		log.Warnf("request offset is less than zero, using fallback '0'")
		input.Offset = 0
	}

	params := GetAnimesParams{
		Query:              input.Query,
		Limit:              input.Limit,
		Offset:             input.Offset,
		IncludeDescription: input.IncludeDescription,
	}

	animes, err := GetAnimesService(context, params)
	if err != nil {
		return response, err
	}

	var result []AnimeTiny

	for _, anime := range animes {
		result = append(result, NewAnimeTiny(&anime))
	}

	response.Body.Animes = result

	return response, nil
}

func GetAnimeInfoController(context context.Context, input *GetAnimeInfoRequest) (*GetAnimeInfoResponse, error) {
	response := &GetAnimeInfoResponse{}

	anime, err := GetAnimeInfoService(context, input.Id)
	if err != nil {
		errId := uuid.New()
		errMsg := "error getting anime info"

		log.Error(errMsg,
			"id", errId,
			"err", err,
		)

		return response, fmt.Errorf("%s, id: %s", errMsg, errId)
	}

	response.Body = anime

	return response, nil
}

func Use(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "remote-search",
		Method:      http.MethodPost,
		Path:        "/remote-search",
		Summary:     "Use anilist to search anime",
	}, RemoteSearchController)

	huma.Register(api, huma.Operation{
		OperationID: "create-anime",
		Method:      http.MethodPost,
		Path:        "/animes",
		Summary:     "Create anime providing anilist ID",
	}, CreateAnimeController)

	huma.Register(api, huma.Operation{
		OperationID: "get-animes",
		Method:      http.MethodGet,
		Path:        "/animes",
		Summary:     "Get anime list",
	}, GetAnimesController)

	huma.Register(api, huma.Operation{
		OperationID: "get-anime",
		Method:      http.MethodGet,
		Path:        "/animes/{id}",
		Summary:     "Get anime info",
	}, GetAnimeInfoController)
}
