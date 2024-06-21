package animes

import (
	"anime-list/httperr"
	"anime-list/services/anilist"
	"anime-list/services/animes/types"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateRequest struct {
	Body struct {
		Id int64 `json:"id"`
	}
}

type CreateResponse struct {
	Body types.Anime
}

func (c *AnimesController) CreateAnime(ctx context.Context, input *CreateRequest) (*CreateResponse, error) {

	resp := &CreateResponse{}

	localId, err := c.repository.TranslateAnilistId(ctx, input.Body.Id)
	if err != nil {
		return nil, httperr.New(http.StatusBadGateway, "error checking if anime already exists", err)
	}

	if localId != 0 {
		return nil, huma.Error404NotFound("anime already exists")
	}

	animeInfo, err := anilist.GetAnime(uint(input.Body.Id))
	if err != nil {
		return nil, httperr.New(http.StatusBadGateway, "error getting anime info", err)
	}

	params := TranslateToParams(animeInfo)

	anime, err := c.service.Create(ctx, params)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error adding anime", err)
	}

	resp.Body = anime

	return resp, nil
}
