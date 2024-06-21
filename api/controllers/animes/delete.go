package animes

import (
	"anime-list/httperr"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteRequest struct {
	Id int64 `path:"id"`
}

type DeleteResponse struct{}

func (c *AnimesController) DeleteAnime(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {

	if request.Id <= 0 {
		return nil, huma.Error400BadRequest("invalid anime id")
	}

	anime, err := c.repository.Get(ctx, request.Id)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting anime", err)
	}

	if anime.Id == 0 {
		return nil, huma.Error404NotFound("anime not found")
	}

	if err := c.repository.Delete(ctx, request.Id); err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error deleting anime", err)
	}

	return &DeleteResponse{}, nil
}
