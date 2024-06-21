package covers

import (
	"anime-list/database"
	"anime-list/repositories/covers/models"
	"context"
)

type CoversRepositoryInterface interface {
	Create(ctx context.Context, params CreateParams) (models.CoverImage, error)
	Get(ctx context.Context, id int64) (models.CoverImage, error)
	GetData(ctx context.Context, animeId int64) (models.CoverData, error)
	Delete(ctx context.Context, id int64) error
}

type CoversRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *CoversRepository {
	return &CoversRepository{db}
}
