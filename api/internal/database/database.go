package database

import (
	"anime-list/internal/database/queries"
	"anime-list/internal/env"
	"database/sql"

	_ "github.com/tursodatabase/go-libsql"
)

var Db *sql.DB
var Query *queries.Queries

func Init() (*sql.DB, error) {

	db, err := sql.Open("libsql", env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = queries.New(db)

	return db, nil
}
