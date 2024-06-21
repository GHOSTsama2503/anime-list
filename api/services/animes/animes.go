package animes

import (
	repo "anime-list/repositories/animes"
	"anime-list/repositories/covers"
	"anime-list/repositories/genres"
	"anime-list/repositories/studios"
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
