package controllers

import (
	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/anime-list/controllers/animes"
	"github.com/ghostsama2503/anime-list/controllers/auth"
	"github.com/ghostsama2503/anime-list/controllers/docs"
	"github.com/ghostsama2503/anime-list/controllers/healthcheck"
	"github.com/ghostsama2503/anime-list/database"
	"github.com/go-chi/chi/v5"

	"github.com/danielgtaylor/huma/v2"
)

func Init(api huma.API, router *chi.Mux, db database.DBTX) {

	if !config.Env.IsProduction {
		router.Get("/docs", docs.ApiDocsHandler)
	}

	auth.Init(api, db)
	animes.Use(api)
	healthcheck.Use(api)
}
