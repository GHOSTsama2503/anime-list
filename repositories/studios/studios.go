package studios

import (
	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/repositories/studios/models"
	"context"
)

type StudiosRepositoryInterface interface {
	Create(ctx context.Context, name string) (models.Studio, error)
	Get(ctx context.Context, id int64) (models.Studio, error)
	GetNames(ctx context.Context, animeId int64) ([]string, error)
	Search(ctx context.Context, name string) (models.Studio, error)
	Update(ctx context.Context, params UpdateParams) (models.Studio, error)
	Delete(ctx context.Context, id int64) error
}

type StudiosRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *StudiosRepository {
	return &StudiosRepository{db}
}
