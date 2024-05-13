package database

import (
	"anime-list/internal/database/queries"
	"anime-list/internal/env"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var Db *sql.DB
var Query *queries.Queries

func Init() (*sql.DB, error) {

	db, err := sql.Open("pgx", env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = queries.New(db)

	return db, nil
}
