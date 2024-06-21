package animes

import (
	"anime-list/services/animes/types"
	"context"
)

type CreateAnimeParams struct {
	IdAl        int64
	Title       types.Title
	Format      string
	Status      string
	Description string
	StartDate   types.Date
	EndDate     types.Date
	Season      string
	SeasonYear  int64
	Episodes    int64
	Duration    int64
	CoverImage  types.CoverImage
	BannerImage string
	Genres      []string
	Studios     []string
}

func (s *AnimesService) Create(ctx context.Context, params CreateAnimeParams) (types.Anime, error) {

	anime := types.Anime{}
	createParams := NewCreateParams(params)

	model, err := s.repository.Create(ctx, createParams)
	if err != nil {
		return anime, err
	}

	anime = NewAnime(model)

	return anime, nil
}
