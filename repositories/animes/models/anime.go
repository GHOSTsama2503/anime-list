package models

import "database/sql"

type Anime struct {
	Id           int64
	IdAl         int64
	TitleRomaji  string
	TitleNative  sql.NullString
	TitleEnglish sql.NullString
	Format       string
	Status       string
	Description  string
	StartDate    string
	EndDate      string
	Season       string
	SeasonYear   sql.NullInt64
	Episodes     int64
	Duration     int64
	BannerImage  sql.NullString
	StImage      string
}
