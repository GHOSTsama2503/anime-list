package animes

import (
	repo "github.com/ghostsama2503/anime-list/api/repositories/animes"
	"github.com/ghostsama2503/anime-list/api/repositories/covers"
	"github.com/ghostsama2503/anime-list/api/repositories/genres"
	"github.com/ghostsama2503/anime-list/api/repositories/studios"
)

type AnimesService struct {
	repository        repo.AnimesRepositoryInterface
	coversRepository  covers.CoversRepositoryInterface
	genresRepository  genres.GenresRepositoryInterface
	studiosRepository studios.StudiosRepositoryInterface
}

func New(repo repo.AnimesRepositoryInterface, coversRepo covers.CoversRepositoryInterface, genresRepo genres.GenresRepositoryInterface, studiosRepo studios.StudiosRepositoryInterface) *AnimesService {
	return &AnimesService{
		repository:        repo,
		coversRepository:  coversRepo,
		genresRepository:  genresRepo,
		studiosRepository: studiosRepo,
	}
}
