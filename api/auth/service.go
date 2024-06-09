package auth

import (
	"anime-list/database"
	"anime-list/database/queries"
	"context"
	"errors"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func AdminExists(ctx context.Context) (bool, error) {

	var exists bool

	if _, err := database.CheckConnection(); err != nil {
		log.Error(err)
		return exists, errors.New("can not connect to database")
	}

	i, err := database.Query.AdminExists(ctx)
	if err != nil {
		return exists, err
	}

	if i >= 1 {
		exists = true
	}

	return exists, nil
}

func CreateUser(ctx context.Context, params CreateUserParams) (User, error) {

	var user User

	if _, err := database.CheckConnection(); err != nil {
		log.Error(err)
		return user, errors.New("can not connect to database")
	}

	var results []queries.User
	results, err := database.Query.GetUserByUsername(ctx, params.Username)
	if err != nil {
		return user, err
	}

	if len(results) >= 1 {
		err = errors.New("username already exists")
		return user, err
	}

	passwdHash, err := CreatePasswordHash(params.Password)
	if err != nil {
		return user, err
	}

	model := params.GetModel()
	model.Password = passwdHash

	newUser := queries.CreateUserParams(model)

	u, err := database.Query.CreateUser(ctx, newUser)
	if err != nil {
		return user, err
	}

	user.New(u)

	return user, nil
}

func CheckUser(ctx context.Context, username string, password string) (User, error) {

	var user User

	if _, err := database.CheckConnection(); err != nil {
		log.Error(err)
		return user, errors.New("can not connect to database")
	}

	results, err := database.Query.GetUserByUsername(ctx, username)
	if err != nil {
		return user, err
	}

	if len(results) != 1 {
		return user, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(results[0].Password), []byte(password))
	if err != nil {
		return user, err
	}

	user.New(results[0])

	return user, nil
}
