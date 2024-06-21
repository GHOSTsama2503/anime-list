package animes

import (
	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/services/animes/types"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetRequest struct {
	Id int64 `path:"id"`
}

type GetResponse struct {
	Body types.Anime
}

func (c *AnimesController) GetAnime(ctx context.Context, request *GetRequest) (*GetResponse, error) {

	anime, err := c.service.Get(ctx, request.Id)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting anime", err)
	}

	if anime.Id == 0 {
		return nil, huma.Error404NotFound("anime not found")
	}

	resp := &GetResponse{}
	resp.Body = anime

	return resp, nil
}
