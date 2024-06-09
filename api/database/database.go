package database

import (
	"anime-list/common/config"
	"anime-list/database/queries"
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

var Db *sql.DB
var Query *queries.Queries
var connected bool

func Init() (*sql.DB, error) {

	db, err := sql.Open("libsql", config.Env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = queries.New(db)
	connected = true

	return db, nil
}

func CheckConnection() (*sql.DB, error) {
	if !connected {
		if err := config.LoadEnv("../.env"); err != nil {
			return Db, err
		}

		return Init()
	}

	return Db, nil
}
