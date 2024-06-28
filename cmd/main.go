package main

import (
	"fmt"
	"net/http"

	"github.com/ghostsama2503/anime-list/common"
	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/anime-list/common/logging"
	"github.com/ghostsama2503/anime-list/controllers"
	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/middlewares"

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
	router.Use(middleware.Compress(5))
	router.Use(middleware.Recoverer)

	apiConfig := huma.DefaultConfig("My Anime List", common.Version())
	apiConfig.DocsPath = ""
	apiConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"cookie": {
			Type: "http",
			Name: "session",
			In:   "cookie",
			Flows: &huma.OAuthFlows{
				Implicit: &huma.OAuthFlow{
					Scopes: map[string]string{"default": "default scope"},
				},
			},
		},
	}

	api := humachi.New(router, apiConfig)

	middlewares.Init(api, router, database.Db)
	controllers.Init(api, router, database.Db)

	log.Info("server running! 🦊", "port", config.Env.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Env.Port), router)
}
