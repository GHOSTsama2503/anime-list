package controllers

import (
	"github.com/ghostsama2503/anime-list/api/controllers/animes"
	"github.com/ghostsama2503/anime-list/api/controllers/healthcheck"
	"github.com/go-chi/chi/v5"

	"github.com/danielgtaylor/huma/v2"
)

func Init(api huma.API, router *chi.Mux) {
	animes.Use(api)
	healthcheck.Use(api)
}
