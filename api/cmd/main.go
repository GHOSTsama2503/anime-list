package main

import (
	"anime-list/internal"
	"anime-list/internal/animes"
	"anime-list/internal/env"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	env.Load()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	if env.IsDevelopment() {
		router.Get("/docs", internal.ApiDocsHandler)
	}

	config := huma.DefaultConfig("My Anime List", "1.0.0")
	config.DocsPath = ""

	api := humachi.New(router, config)

	animes.Use(api)

	address := fmt.Sprintf(":%d", env.PORT)
	http.ListenAndServe(address, router)
}
