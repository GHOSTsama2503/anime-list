package animes

import (
	"anime-list/internal/anilist"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
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

	_ = result

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
		Summary:     "Create",
	}, CreateAnimeController)
}
