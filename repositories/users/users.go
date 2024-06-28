package users

import (
	"context"

	"github.com/ghostsama2503/anime-list/database"
	"github.com/ghostsama2503/anime-list/repositories/users/models"
)

type UsersRepositoryInterface interface {
	Create(ctx context.Context, params CreateParams) (models.User, error)
	Get(ctx context.Context, parms GetParams) (models.User, error)
	Delete(ctx context.Context, id int64) error

	CheckIfAdminExists(ctx context.Context) (bool, error)

	GetPassword(ctx context.Context, id int64) (string, error)
	CheckPassword(ctx context.Context, params CheckPasswordParams) (bool, error)
}

type UsersRepository struct {
	db database.DBTX
}

func NewUsersRepository(db database.DBTX) *UsersRepository {
	return &UsersRepository{db}
}
