package studios_test

import (
	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/repositories/studios"
	"context"
	"database/sql"
	"testing"
)

func TestStudiosRepository(t *testing.T) {

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

	repository := studios.New(tx)

	newStudio, err := repository.Create(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}

	studio, err := repository.Get(ctx, newStudio.Id)
	if err != nil {
		t.Fatal(err)
	}

	studioResult, err := repository.Search(ctx, studio.Name)
	if err != nil {
		t.Fatal(err)
	}

	updateParams := studios.UpdateParams{
		Id:   studioResult.Id,
		Name: "bar",
	}

	updatedStudio, err := repository.Update(ctx, updateParams)
	if err != nil {
		t.Fatal(err)
	}

	if err := repository.Delete(ctx, updatedStudio.Id); err != nil {
		t.Fatal(err)
	}
}
