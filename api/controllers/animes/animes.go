package animes

import (
	"anime-list/database"
	repository "anime-list/repositories/animes"
	"anime-list/repositories/covers"
	"anime-list/repositories/genres"
	"anime-list/repositories/studios"
	service "anime-list/services/animes"
	"net/http"

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
