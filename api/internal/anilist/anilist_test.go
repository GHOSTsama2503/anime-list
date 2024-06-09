package anilist

import (
	"strings"
	"testing"
)

func TestSearchAnime(t *testing.T) {
	const page uint = 0
	const perPage uint = 10

	results, err := SearchAnime("Naruto", page, perPage)
	if err != nil {
		t.Fatal(err)
	}

	if !(len(results) >= 3) {
		t.Fatal("unexpected lenght of SearchAnime results")
	}

	if !(results[0].Id > 0) {
		t.Fatal("anime with ID that is not greater than 0")
	}

	if !strings.HasPrefix(strings.ToLower(results[1].Image.Medium), "https://") {
		t.Fatal("anime image url that doesn't start with https://")
	}
}

func TestGetAnime(t *testing.T) {
	const id uint = 20

	result, err := GetAnime(id)
	if err != nil {
		t.Fatal(err)
	}

	if !(result.Id > 0) {
		t.Fatal("anime with ID that is not greater than 0")
	}
}
