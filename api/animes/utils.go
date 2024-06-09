package animes

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: str != ""}
}

func NullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

func StImage(id int) string {
	return fmt.Sprintf("https://img.anili.st/media/%d", id)
}

func FuzzyDateToTime(d FuzzyDate) time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, &time.Location{})
}

func FuzzyDateToString(d FuzzyDate) string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

func SFuzzyDateToTime(s string) (t time.Time, err error) {
	t, err = time.Parse(time.DateOnly, s)
	return
}

func StringToFuzzyDate(s string) (d FuzzyDate, err error) {

	errMsg := "can not parse string to fuzzy date"

	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		err = fmt.Errorf("%s, invalid string format: %s", errMsg, s)
		return
	}

	year := parts[0]
	if d.Year, err = strconv.Atoi(year); err != nil {
		err = fmt.Errorf("%s, invalid year: %s", errMsg, year)
		return
	}

	month := parts[1]
	if d.Month, err = strconv.Atoi(month); err != nil {
		err = fmt.Errorf("%s, invalid month: %s", errMsg, month)
		return
	}

	day := parts[2]
	if d.Day, err = strconv.Atoi(day); err != nil {
		err = fmt.Errorf("%s, invalid day: %s", errMsg, day)
		return
	}

	return
}

func DateToTime(date string) (time.Time, error) {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return time.Time{}, nil
	}

	return parsedDate, nil
}

func IsValidFormat(format string) bool {
	// TODO: missing validator
	return true
}

func NewAnimeTiny(anime *Anime) (result AnimeTiny) {
	result.Id = anime.Id
	result.Title = anime.Title.Romaji
	result.Description = anime.Description
	result.ImageId = anime.IdAl
	return
}
