package animes

import "anime-list/internal/database/queries"

type Season string

const (
	SeasonFall   Season = "FALL"
	SeasonSpring Season = "SPRING"
	SeasonSummer Season = "SUMMER"
	SeasonWinter Season = "WINTER"
	SeasonUnknow Season = "UNKNOWN"
)

type Anime struct {
	Id          int        `json:"id"`
	IdAl        int        `json:"id_al"`
	Title       Title      `json:"title"`
	Format      string     `json:"format"`
	Status      string     `json:"status"`
	Description string     `json:"description"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
	Season      string     `json:"season"`
	SeasonYear  int        `json:"season_year"`
	Episodes    int        `json:"episodes"`
	Duration    int        `json:"duration"`
	Genres      []string   `json:"genres"`
	Studios     []string   `json:"studios"`
	CoverImage  CoverImage `json:"cover_image"`
	BannerImage string     `json:"banner_image"`
	StImage     string     `json:"st_image"`
}

func (anime *Anime) New(model queries.Anime, genres, studios []string, cover CoverImage) {
	anime.Id = int(model.ID)
	anime.IdAl = int(model.IDAl)
	anime.Title = Title{
		Romaji:  model.TitleRomaji,
		Native:  model.TitleNative.String,
		English: model.TitleEnglish.String,
	}
	anime.Format = model.Format
	anime.Status = model.Status
	anime.Description = model.Description
	anime.StartDate = model.StartDate
	anime.EndDate = model.EndDate
	anime.Season = model.Season
	anime.SeasonYear = int(model.SeasonYear.Int64)
	anime.Episodes = int(model.Episodes)
	anime.Duration = int(model.Duration)
	anime.Genres = genres
	anime.Studios = studios
	anime.BannerImage = model.BannerImage.String
	anime.StImage = model.StImage
}

type Title struct {
	Romaji  string `json:"romaji"`
	Native  string `json:"native"`
	English string `json:"english"`
}

type FuzzyDate struct {
	Year  int `json:"year" binding:"required"`
	Month int `json:"month" binding:"required"`
	Day   int `json:"day" binding:"required"`
}

type Trailer struct{}

type CoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Color      string `json:"color"`
}
