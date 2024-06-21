package controllers

import (
	"github.com/ghostsama2503/anime-list/api/common/config"
	"github.com/ghostsama2503/anime-list/api/controllers/animes"
	"github.com/ghostsama2503/anime-list/api/controllers/docs"
	"github.com/ghostsama2503/anime-list/api/controllers/healthcheck"
	"github.com/go-chi/chi/v5"

	"github.com/danielgtaylor/huma/v2"
)

func Init(api huma.API, router *chi.Mux) {

	if !config.Env.IsProduction {
		router.Get("/docs", docs.ApiDocsHandler)
	}

	animes.Use(api)
	healthcheck.Use(api)
}
