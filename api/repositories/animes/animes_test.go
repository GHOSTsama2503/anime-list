package animes_test

import (
	"anime-list/database"
	"anime-list/repositories/animes"
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestAnimesRepository(t *testing.T) {

	db, err := database.CheckConnection()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		t.Fatal(err)
	}

	defer tx.Rollback()

	repository := animes.New(tx)

	createParams := animes.CreateParams{
		TitleRomaji: "foo",
	}

	newAnime, err := repository.Create(ctx, createParams)
	if err != nil {
		t.Fatal(err)
	}

	anime, err := repository.Get(ctx, newAnime.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err := repository.Delete(ctx, anime.Id); err != nil {
		t.Fatal(err)
	}

	getListParams := animes.GetListParams{
		Limit: 10,
	}

	animeResult, err := repository.GetList(ctx, getListParams)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range animeResult {
		fmt.Println(v.Title)
	}
}
