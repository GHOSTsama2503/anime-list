package animes

type RemoteSearchRequest struct {
	Body struct {
		Title string `json:"title"`
	}
}

type RemoteSearchResponse struct {
	Body []RemoteSearchResult `json:"result"`
}

type RemoteSearchResult struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type CreateAnimeRequest struct {
	Body struct {
		Id            uint   `json:"id"`
		Group         string `json:"group,omitempty"`
		GroupPosition int    `json:"group_position,omitempty" required:"false"`
	}
}

type CreateAnimeResponse struct {
	Body struct {
		Ok bool `json:"ok"`
	} `json:"body"`
}

type GetAnimesRequest struct {
	Query              string `query:"q"`
	Offset             int64  `query:"offset"`
	Limit              int64  `query:"limit"`
	IncludeDescription bool   `query:"desc"`
}

type GetAnimesResponse struct {
	Body struct {
		Total  int         `json:"total"`
		Animes []AnimeTiny `json:"animes"`
	}
}

type AnimeTiny struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageId     int    `json:"image"`
}

type GetAnimeInfoRequest struct {
	Id int `path:"id"`
}

type GetAnimeInfoResponse struct {
	Body Anime
}

// type CreateAnimeParams struct {
// 	IdAl            int        `json:"id_al" binding:"required" doc:"Anilist ID" minimum:"1"`
// 	IdMal           int        `json:"id_mal" binding:"required" doc:"MyAnimeList ID" minimum:"1"`
// 	Title           Title      `json:"title" binding:"required"`
// 	Format          string     `json:"format" binding:"required"`
// 	Status          string     `json:"status" binding:"required"`
// 	Description     string     `json:"description" binding:"required"`
// 	StartDate       string     `json:"start_date" binding:"required" doc:"Date format: YYYY-MM-DD"`
// 	EndDate         string     `json:"end_date" binding:"required" doc:"Date format: YYYY-MM-DD"`
// 	Season          string     `json:"season" binding:"required"`
// 	SeasonYear      int        `json:"season_year" binding:"required"`
// 	Episodes        int        `json:"episodes" binding:"required"`
// 	Duration        int        `json:"duration" binding:"required"`
// 	Genres          []string   `json:"genres" binding:"required"`
// 	Studios         []string   `json:"studios" binding:"required"`
// 	CountryOfOrigin string     `json:"country_of_origin" binding:"required"`
// 	Hashtag         string     `json:"hashtag"`
// 	Trailer         Trailer    `json:"trailer"`
// 	CoverImage      CoverImage `json:"cover_image" binding:"required"`
// 	BannerImage     string     `json:"banner_image" binding:"required"`
// 	Score           int        `json:"score" binding:"required"`
// }
