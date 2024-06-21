package animes

import (
	"github.com/ghostsama2503/anime-list/api/database"
	"github.com/ghostsama2503/anime-list/api/repositories/animes/models"
	"context"
)

type AnimesRepositoryInterface interface {
	Create(ctx context.Context, params CreateParams) (models.Anime, error)
	Get(ctx context.Context, id int64) (models.Anime, error)
	GetList(ctx context.Context, params GetListParams) ([]models.AnimeTiny, error)
	Delete(ctx context.Context, id int64) error

	TranslateAnilistId(ctx context.Context, id int64) (int64, error)
}

type AnimesRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *AnimesRepository {
	return &AnimesRepository{db}
}
