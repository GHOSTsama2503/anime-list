package anilist

type SearchResult struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       struct {
		ExtralLarge string `json:"extraLarge"`
		Large       string `json:"large"`
		Medium      string `json:"medium"`
		Color       string `json:"color"`
	} `json:"image"`
}

type SearchRequest struct {
	Query     string                 `json:"query"`
	Variables SearchRequestVariables `json:"variables"`
}

type SearchRequestVariables struct {
	Search  string `json:"search"`
	Page    uint   `json:"page"`
	PerPage uint   `json:"perPage"`
}

type SearchResponse struct {
	Data struct {
		Page struct {
			PageInfo struct {
				Total       uint `json:"total"`
				CurrentPage uint `json:"currentPage"`
				LastPage    uint `json:"lastPage"`
				HasNextPage bool `json:"hasNextPage"`
				PerPage     uint `json:"perPage"`
			} `json:"pageInfo"`
			Media []struct {
				Id    uint `json:"id"`
				Title struct {
					Romaji string `json:"romaji"`
				} `json:"title"`
				Description string `json:"description"`
				CoverImage  struct {
					ExtralLarge string `json:"extraLarge"`
					Large       string `json:"large"`
					Medium      string `json:"medium"`
					Color       string `json:"color"`
				} `json:"coverImage"`
				BannerImage string `json:"bannerImage"`
			} `json:"media"`
		} `json:"page"`
	} `json:"data"`
}

type GetAnimeRequest struct {
	Query     string `json:"query"`
	Variables struct {
		Id uint `json:"id"`
	} `json:"variables"`
}

type GetAnimeResponse struct {
	Data struct {
		Media Anime `json:"Media"`
	} `json:"data"`
}

type Anime struct {
	Id    uint `json:"id"`
	Title struct {
		Romaji  string `json:"romaji"`
		Native  string `json:"native"`
		English string `json:"english"`
	} `json:"title"`
	Format      string     `json:"format"`
	Status      string     `json:"status"`
	Description string     `json:"description"`
	StartDate   FuzzyDate  `json:"startDate"`
	EndDate     FuzzyDate  `json:"endDate"`
	Season      string     `json:"season"`
	Episodes    int        `json:"episodes"`
	Duration    int        `json:"duration"`
	SeasonYear  int        `json:"seasonYear"`
	CoverImage  CoverImage `json:"coverImage"`
	BannerImage string     `json:"bannerImage"`
	Genres      []string   `json:"genres"`
	Studios     struct {
		Nodes []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"studios"`
}

type FuzzyDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type CoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Color      string `json:"color"`
}
