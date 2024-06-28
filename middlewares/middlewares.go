package middlewares

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/database"
	sessionsRepo "github.com/ghostsama2503/anime-list/repositories/sessions"
	"github.com/ghostsama2503/anime-list/services/sessions"

	"github.com/go-chi/chi/v5"
)

type HumaMiddleware = func(ctx huma.Context, next func(huma.Context))

func Init(api huma.API, router *chi.Mux, db database.DBTX) {

	sessionsRepository := sessionsRepo.New(db)
	sessionsService := sessions.New(sessionsRepository)

	api.UseMiddleware(Auth(api, sessionsService))
}
