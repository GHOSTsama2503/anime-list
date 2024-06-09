package animes

type CreateAnimeParams struct {
	IdAl          int
	Title         Title
	Format        string
	Status        string
	Description   string
	StartDate     FuzzyDate
	EndDate       FuzzyDate
	Season        Season
	SeasonYear    int
	Episodes      int
	Duration      int
	CoverImage    CoverImage
	BannerImage   string
	Genres        []string
	Studios       []string
	Group         string
	GroupPosition int
}