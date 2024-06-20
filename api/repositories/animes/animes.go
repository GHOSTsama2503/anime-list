package animes

import (
	"anime-list/database"
	"anime-list/repositories/animes/models"
	"context"
)

type AnimesRepositoryInterface interface {
	Create(ctx context.Context, params CreateParams) (models.Anime, error)
	Get(ctx context.Context, id int64) (models.Anime, error)
	GetList(ctx context.Context, params GetListParams) ([]models.AnimeTiny, error)
	Update(ctx context.Context) (models.Anime, error)
	Delete(ctx context.Context, id int64) error
}

type AnimesRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *AnimesRepository {
	return &AnimesRepository{db}
}
