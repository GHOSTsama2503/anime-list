package main

import (
	"anime-list/animes"
	"anime-list/auth"
	"anime-list/common"
	"anime-list/common/config"
	"anime-list/common/logging"
	"anime-list/database"
	"anime-list/docs"
	"anime-list/healthcheck"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		panic(fmt.Sprintf("error loading environment: %v", err))
	}

	logging.SetupLogger()

	_, err := database.Init()
	if err != nil {
		log.Fatal("can not initialize database", "err", err)
	}

	defer database.Db.Close()

	router := chi.NewRouter()
	router.Use(logging.LoggerMidleware)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	if !config.Env.IsProduction {
		router.Get("/docs", docs.ApiDocsHandler)
	}

	apiConfig := huma.DefaultConfig("My Anime List", common.Version())
	apiConfig.DocsPath = ""

	api := humachi.New(router, apiConfig)

	// middlewares
	api.UseMiddleware(auth.NewAuthMiddleware(api))

	// routes
	auth.Use(api)
	animes.Use(api)
	healthcheck.Use(api)

	log.Info("server running! 🦊", "port", config.Env.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Env.Port), router)
}
