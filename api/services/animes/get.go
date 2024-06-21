package animes

import (
	"github.com/ghostsama2503/anime-list/api/services/animes/types"
	"context"
)

func (s *AnimesService) Get(ctx context.Context, id int64) (types.Anime, error) {

	var anime types.Anime

	model, err := s.repository.Get(ctx, id)
	if err != nil {
		return anime, err
	}

	coverImage, err := s.coversRepository.GetData(ctx, id)
	if err != nil {
		return anime, err
	}

	genres, err := s.genresRepository.GetNames(ctx, id)
	if err != nil {
		return anime, err
	}

	studios, err := s.studiosRepository.GetNames(ctx, id)
	if err != nil {
		return anime, err
	}

	anime = NewAnime(model)
	anime.CoverImage = coverImage
	anime.Genres = genres
	anime.Studios = studios

	return anime, nil
}
