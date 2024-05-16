// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"database/sql"
)

type Anime struct {
	ID            int64
	IDAl          int64
	TitleRomaji   string
	TitleNative   sql.NullString
	TitleEnglish  sql.NullString
	Format        string
	Status        string
	Description   string
	StartDate     string
	EndDate       string
	Season        string
	SeasonYear    sql.NullInt64
	Episodes      int64
	Duration      int64
	BannerImage   sql.NullString
	StImage       string
	GroupPosition sql.NullInt64
}

type AnimeGenre struct {
	AnimeID int64
	GenreID int64
}

type AnimeGroup struct {
	AnimeID int64
	GroupID int64
}

type AnimeStudio struct {
	AnimeID  int64
	StudioID int64
}

type CoverImage struct {
	ID         int64
	AnimeID    int64
	ExtraLarge sql.NullString
	Large      sql.NullString
	Medium     sql.NullString
	Color      sql.NullString
}

type Genre struct {
	ID   int64
	Name string
}

type Group struct {
	ID   int64
	Name string
}

type Studio struct {
	ID   int64
	Name string
}

type Synonym struct {
	ID      int64
	AnimeID int64
	Name    string
}
