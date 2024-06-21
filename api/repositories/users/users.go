package users

import (
	"github.com/ghostsama2503/anime-list/api/database"
	"github.com/ghostsama2503/anime-list/api/repositories/users/models"
	"context"
)

type UsersRepositoryInterface interface {
	Create(ctx context.Context, params CreateParams) (models.User, error)
	GetUserById(ctx context.Context, id int64) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	DeleteUser(ctx context.Context, id int64) error

	CheckIfAdminExists() (bool, error)
	CheckUserPassword(ctx context.Context, params CheckUserPasswordParams) (bool, error)
}

type UsersRepository struct {
	db database.DBTX
}

func NewUsersRepository(db database.DBTX) *UsersRepository {
	return &UsersRepository{db}
}
