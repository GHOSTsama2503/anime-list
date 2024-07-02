package animes

import (
	"context"
	"net/http"

	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/services/animes"
	"github.com/ghostsama2503/anime-list/services/animes/types"
)

type GetListRequest struct {
	Limit               int64  `query:"limit"`
	Offset              int64  `query:"offset"`
	Query               string `query:"q"`
	IncludeDescriptions bool   `query:"desc"`
}

type GetListResponse struct {
	Body struct {
		Animes []types.AnimeTiny `json:"animes"`
	}
}

func (c *AnimesController) GetAnimeList(ctx context.Context, request *GetListRequest) (*GetListResponse, error) {

	params := animes.GetListParams{
		Query:               request.Query,
		Limit:               request.Limit,
		Offset:              request.Offset,
		IncludeDescriptions: request.IncludeDescriptions,
	}

	animesList, err := c.service.GetList(ctx, params)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting anime list", err)
	}

	resp := &GetListResponse{}
	resp.Body.Animes = animesList

	return resp, nil
}
