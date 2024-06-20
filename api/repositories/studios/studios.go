package studios

import (
	"anime-list/database"
	"anime-list/repositories/studios/models"
	"context"
)

type StudiosRepositoryInterface interface {
	Create(ctx context.Context, name string) (models.Studio, error)
	Get(ctx context.Context, id int64) (models.Studio, error)
	Search(ctx context.Context, name string) (models.Studio, error)
	Update(ctx context.Context, params UpdateParams) (models.Studio, error)
	Delete(ctx context.Context, id int64) (models.Studio, error)
}

type StudiosRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *StudiosRepository {
	return &StudiosRepository{db}
}
