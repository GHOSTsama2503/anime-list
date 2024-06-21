package types

import "anime-list/repositories/covers/models"

type Anime struct {
	Id          int64            `json:"id"`
	IdAl        int64            `json:"idAl"`
	Title       Title            `json:"title"`
	Format      string           `json:"format"`
	Status      string           `json:"status"`
	Description string           `json:"description"`
	StartDate   string           `json:"startDate"`
	EndDate     string           `json:"endDate"`
	Season      string           `json:"season"`
	SeasonYear  int64            `json:"seasonYear"`
	Episodes    int64            `json:"episodes"`
	Duration    int64            `json:"duration"`
	BannerImage string           `json:"bannerImage"`
	StImage     string           `json:"stImage"`
	CoverImage  models.CoverData `json:"coverImage"`
	Genres      []string         `json:"genres"`
	Studios     []string         `json:"studios"`
}
