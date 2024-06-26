package genres_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/repositories/genres"
)

func TestGenresRepository(t *testing.T) {
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

	repository := genres.New(tx)

	newGenre, err := repository.Create(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}

	genre, err := repository.Get(ctx, newGenre.Id)
	if err != nil {
		t.Fatal(err)
	}

	genreResult, err := repository.Search(ctx, genre.Name)
	if err != nil {
		t.Fatal(err)
	}

	updateParams := genres.UpdateParams{
		Id:   genreResult.Id,
		Name: "bar",
	}

	updatedGenre, err := repository.Update(ctx, updateParams)
	if err != nil {
		t.Fatal(err)
	}

	if err := repository.Delete(ctx, updatedGenre.Id); err != nil {
		t.Fatal(err)
	}
}
