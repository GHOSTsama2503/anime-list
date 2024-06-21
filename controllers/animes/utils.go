package animes

import (
	"github.com/ghostsama2503/anime-list/api/services/anilist"
	"github.com/ghostsama2503/anime-list/api/services/animes"
	"github.com/ghostsama2503/anime-list/api/services/animes/types"
)

func TranslateToParams(anime anilist.Anime) animes.CreateAnimeParams {

	var studios []string
	for _, studio := range anime.Studios.Nodes {
		studios = append(studios, studio.Name)
	}

	return animes.CreateAnimeParams{
		IdAl:        int64(anime.Id),
		Title:       anime.Title,
		Format:      anime.Format,
		Status:      anime.Status,
		Description: anime.Description,
		StartDate:   types.Date(anime.StartDate),
		EndDate:     types.Date(anime.EndDate),
		Season:      anime.Season,
		SeasonYear:  int64(anime.SeasonYear),
		Episodes:    int64(anime.Episodes),
		Duration:    int64(anime.Duration),
		CoverImage:  types.CoverImage(anime.CoverImage),
		BannerImage: anime.BannerImage,
		Genres:      anime.Genres,
		Studios:     studios,
	}
}
