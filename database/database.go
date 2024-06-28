package database

import (
	"database/sql"

	"github.com/ghostsama2503/anime-list/common/config"
	_ "modernc.org/sqlite"
)

var Db *sql.DB
var Query *Queries
var connected bool

// Initialize database connection
func Init() (*sql.DB, error) {

	db, err := sql.Open("sqlite", config.Env.DatabaseUrl)
	if err != nil {
		return db, err
	}

	Db = db
	Query = New(db)
	connected = true

	return db, nil
}

// Returns current database connection if exists,
// else creates a new connection and return it.
func CheckConnection() (*sql.DB, error) {

	if !connected {
		if err := config.LoadEnv(); err != nil {
			return Db, err
		}

		return Init()
	}

	return Db, nil
}
