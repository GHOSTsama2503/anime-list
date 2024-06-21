package animes

import (
	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/repositories/animes"
	"github.com/ghostsama2503/anime-list/repositories/animes/models"
	"context"
	"net/http"
)

type GetListRequest struct {
	Limit               int64  `query:"limit"`
	Offset              int64  `query:"offset"`
	Query               string `query:"q"`
	IncludeDescriptions bool   `query:"desc"`
}

type GetListResponse struct {
	Body struct {
		Animes []models.AnimeTiny `json:"animes"`
	}
}

func (c *AnimesController) GetAnimeList(ctx context.Context, input *GetListRequest) (*GetListResponse, error) {

	limit := input.Limit
	if limit > config.Env.MaxLimit {
		limit = config.Env.DefaultLimit
	}

	offset := input.Offset
	if offset < 0 {
		offset = 0
	}

	params := animes.GetListParams{
		Limit:               limit,
		Offset:              offset,
		Query:               input.Query,
		IncludeDescriptions: input.IncludeDescriptions,
	}

	animes, err := c.repository.GetList(ctx, params)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting anime list", err)
	}

	resp := &GetListResponse{}
	resp.Body.Animes = animes

	return resp, nil
}
