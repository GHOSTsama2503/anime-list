package animes

import (
	"net/http"

	"github.com/ghostsama2503/anime-list/database"
	repository "github.com/ghostsama2503/anime-list/repositories/animes"
	"github.com/ghostsama2503/anime-list/repositories/covers"
	"github.com/ghostsama2503/anime-list/repositories/genres"
	"github.com/ghostsama2503/anime-list/repositories/studios"
	service "github.com/ghostsama2503/anime-list/services/animes"

	"github.com/danielgtaylor/huma/v2"
)

type AnimesController struct {
	service    *service.AnimesService
	repository repository.AnimesRepositoryInterface
}

func Use(api huma.API) error {

	db, err := database.CheckConnection()
	if err != nil {
		return err
	}

	repo := repository.New(db)
	srv := service.New(repo, covers.New(db), genres.New(db), studios.New(db))

	controller := AnimesController{
		service:    srv,
		repository: repo,
	}

	var group = "animes"

	huma.Register(api, huma.Operation{
		Method:  http.MethodPost,
		Path:    "/animes",
		Summary: "create anime",
		Tags:    []string{group},
	}, controller.CreateAnime)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/animes",
		Summary: "get anime list",
		Tags:    []string{group},
	}, controller.GetAnimeList)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/animes/{id}",
		Summary: "get anime info",
		Tags:    []string{group},
	}, controller.GetAnime)

	huma.Register(api, huma.Operation{
		Method:  http.MethodDelete,
		Path:    "/animes/{id}",
		Summary: "delete anime",
		Tags:    []string{group},
	}, controller.DeleteAnime)

	return nil
}
