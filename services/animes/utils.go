package animes

import (
	"fmt"

	r "github.com/ghostsama2503/anime-list/repositories/animes"
	"github.com/ghostsama2503/anime-list/repositories/animes/models"
	coversModels "github.com/ghostsama2503/anime-list/repositories/covers/models"
	"github.com/ghostsama2503/anime-list/services"
	"github.com/ghostsama2503/anime-list/services/animes/types"
)

func DateToString(d types.Date) string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

func NewCreateParams(params CreateAnimeParams) r.CreateParams {
	return r.CreateParams{
		IdAl:         params.IdAl,
		TitleRomaji:  params.Title.Romaji,
		TitleEnglish: services.NullString(params.Title.English),
		TitleNative:  services.NullString(params.Title.Native),
		Format:       params.Format,
		Status:       params.Format,
		Description:  params.Description,
		StartDate:    DateToString(params.StartDate),
		EndDate:      DateToString(params.EndDate),
		Season:       params.Season,
		SeasonYear:   services.NullInt64(params.SeasonYear),
	}
}

func NewAnime(model models.Anime) types.Anime {
	return types.Anime{
		Id:          model.Id,
		IdAl:        model.IdAl,
		Title:       types.Title{Romaji: model.TitleRomaji, Native: model.TitleNative.String, English: model.TitleEnglish.String},
		Format:      model.Format,
		Status:      model.Status,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Season:      model.Season,
		SeasonYear:  model.SeasonYear.Int64,
		Episodes:    model.Episodes,
		Duration:    model.Duration,
		BannerImage: model.BannerImage.String,
		StImage:     model.StImage,
	}
}

func NewAnimeTiny(model models.AnimeTiny) types.AnimeTiny {
	return types.AnimeTiny{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
	}
}

func NewCoverImage(model coversModels.CoverImage) types.CoverImage {
	return types.CoverImage{
		ExtraLarge: model.ExtraLarge,
		Large:      model.Large,
		Medium:     model.Medium,
		Color:      model.Color,
	}
}
