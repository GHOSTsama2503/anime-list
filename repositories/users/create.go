package users

import (
	"github.com/ghostsama2503/anime-list/repositories/users/models"
	"context"
)

const create = `
INSERT INTO users (username, password, is_admin)
VALUES (?, ?, ?)
RETURNING id, username, is_admin;
`

type CreateParams struct {
	Username string
	Password string
	IsAdmin  bool
}

func (r *UsersRepository) Create(ctx context.Context, params CreateParams) (models.User, error) {

	var isAdminParam int64 = 0
	if params.IsAdmin {
		isAdminParam = 1
	}

	row := r.db.QueryRowContext(ctx, create, params.Username, params.Password, isAdminParam)

	var user models.User

	var isAdmin int64

	err := row.Scan(&user.Id, &user.Username, &isAdmin)
	if err != nil {
		return user, err
	}

	user.IsAdmin = params.IsAdmin

	return user, nil
}
