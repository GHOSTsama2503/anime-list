package database_test

import (
	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/anime-list/database"
	"testing"
)

func TestInit(t *testing.T) {

	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	var err error
	if _, err = database.Init(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckConnection(t *testing.T) {
	var err error
	if _, err = database.CheckConnection(); err != nil {
		t.Fatal(err)
	}
}
