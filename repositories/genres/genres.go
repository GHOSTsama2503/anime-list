package genres

import (
	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/repositories/genres/models"
	"context"
)

type GenresRepositoryInterface interface {
	Create(ctx context.Context, name string) (models.Genre, error)
	Get(ctx context.Context, id int64) (models.Genre, error)
	GetNames(ctx context.Context, animeId int64) ([]string, error)
	Search(ctx context.Context, name string) (models.Genre, error)
	Update(ctx context.Context, params UpdateParams) (models.Genre, error)
	Delete(ctx context.Context, id int64) error
}

type GenresRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *GenresRepository {
	return &GenresRepository{db}
}
