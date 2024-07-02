package animes

import (
	"context"

	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/anime-list/services/animes/types"
)

type GetListParams struct {
	Query               string
	Limit               int64
	Offset              int64
	IncludeDescriptions bool
}

func (params GetListParams) GetQuery() string {
	return params.Query
}

func (params GetListParams) GetLimit() int64 {
	return params.Limit
}

func (params GetListParams) GetOffset() int64 {
	return params.Offset
}

func (params GetListParams) GetIncludeDescriptions() bool {
	return params.IncludeDescriptions
}

func (service *AnimesService) GetList(ctx context.Context, params GetListParams) ([]types.AnimeTiny, error) {

	limit := params.Limit
	if limit > config.Env.MaxLimit || limit == 0 {
		limit = config.Env.DefaultLimit
	}

	offset := params.Offset
	if offset < 0 {
		offset = 0
	}

	repositoryParams := GetListParams{
		Query:               params.Query,
		Limit:               limit,
		Offset:              offset,
		IncludeDescriptions: params.IncludeDescriptions,
	}

	modelsList, err := service.repository.GetList(ctx, repositoryParams)
	if err != nil {
		return nil, err
	}

	animes := make([]types.AnimeTiny, 0)

	for _, model := range modelsList {
		animes = append(animes, NewAnimeTiny(model))
	}

	return animes, nil
}
