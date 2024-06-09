package main

import (
	"anime-list/common"
	"anime-list/internal"
	"anime-list/internal/animes"
	"anime-list/internal/database"
	"anime-list/internal/env"
	"anime-list/internal/healthcheck"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("error loading environment: %v", err))
	}

	internal.SetupLogger()

	_, err := database.Init()
	if err != nil {
		log.Fatal("can not initialize database", "err", err)
	}

	defer database.Db.Close()

	router := chi.NewRouter()
	router.Use(internal.LoggerMidleware)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	if env.IsDevelopment {
		router.Get("/docs", internal.ApiDocsHandler)
	}

	config := huma.DefaultConfig("My Anime List", common.Version())
	config.DocsPath = ""

	api := humachi.New(router, config)
	animes.Use(api)
	healthcheck.Use(api)

	log.Info("server running! 🦊", "port", env.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", env.Port), router)
}
