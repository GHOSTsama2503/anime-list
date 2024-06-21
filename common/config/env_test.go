package config_test

import (
	"github.com/ghostsama2503/anime-list/api/common/config"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

	if config.Env.IsProduction {
		t.Fatal("unexpected environment")
	}

	if config.Env.DatabaseUrl == "" {
		t.Fatal("database url is empty")
	}
}
