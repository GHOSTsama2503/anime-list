package users_test

import (
	"github.com/ghostsama2503/anime-list/api/database"
	"github.com/ghostsama2503/anime-list/api/repositories/users"
	"context"
	"database/sql"
	"testing"
)

func TestUsersRepository(t *testing.T) {

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

	usersRepository := users.NewUsersRepository(tx)

	const userPassword = "bar"

	createUserParams := users.CreateParams{
		Username: "foo",
		Password: userPassword,
	}

	newUser, err := usersRepository.Create(ctx, createUserParams)
	if err != nil {
		t.Fatal(err)
	}

	getParams := users.GetParams{
		Id: newUser.Id,
	}

	user, err := usersRepository.Get(ctx, getParams)
	if err != nil {
		t.Fatal(err)
	}

	adminExists, err := usersRepository.CheckIfAdminExists(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_ = adminExists

	checkUserPasswordParams := users.CheckUserPasswordParams{
		Username: user.Username,
		Password: userPassword,
	}

	validPasswordIsValid := usersRepository.CheckPassword(ctx, checkUserPasswordParams)
	if err != nil {
		t.Fatal(err)
	}

	if !validPasswordIsValid {
		t.Fatal("invalid password verification result, expected true")
	}

	invalidPasswordParams := users.CheckUserPasswordParams{
		Username: user.Username,
		Password: userPassword + "invalid",
	}

	invalidPasswordIsValid := usersRepository.CheckPassword(ctx, invalidPasswordParams)
	if err != nil {
		t.Fatal(err)
	}

	if invalidPasswordIsValid {
		t.Fatalf("invalid password verification result, expected false")
	}

	if err := usersRepository.Delete(ctx, user.Id); err != nil {
		t.Fatal(err)
	}
}
