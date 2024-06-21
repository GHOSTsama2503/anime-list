package database

import (
	"github.com/ghostsama2503/anime-list/common/config"
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

var Db *sql.DB
var Query *Queries
var connected bool

func Init() (*sql.DB, error) {

	db, err := sql.Open("libsql", config.Env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = New(db)
	connected = true

	return db, nil
}

func CheckConnection() (*sql.DB, error) {
	if !connected {
		if err := config.LoadEnv(); err != nil {
			return Db, err
		}

		return Init()
	}

	return Db, nil
}
