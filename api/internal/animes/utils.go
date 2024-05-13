package animes

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func NullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: str != ""}
}

func NullInt16(i int16) sql.NullInt16 {
	return sql.NullInt16{Int16: i, Valid: i != 0}
}

func StImage(id int) string {
	return fmt.Sprintf("https://img.anili.st/media/%d", id)
}

func FuzzyDateToTime(d FuzzyDate) time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, &time.Location{})
}

func DateToTime(date string) (time.Time, error) {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return time.Time{}, nil
	}

	return parsedDate, nil
}

func IsValidSeason(season string) bool {
	season = strings.ToUpper(season)

	return (season == SeasonFall ||
		season == SeasonSpring ||
		season == SeasonSummer ||
		season == SeasonWinter ||
		season == SeasonUnknow)
}

func IsValidFormat(format string) bool {
	// TODO: missing validator
	return true
}
