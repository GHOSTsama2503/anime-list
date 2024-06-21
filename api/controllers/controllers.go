package controllers

import (
	"anime-list/controllers/animes"
	"anime-list/controllers/healthcheck"

	"github.com/danielgtaylor/huma/v2"
)

func Init(api huma.API) {
	animes.Use(api)
	healthcheck.Use(api)
}
