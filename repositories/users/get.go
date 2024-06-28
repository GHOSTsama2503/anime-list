package users

import (
	"context"

	"github.com/ghostsama2503/anime-list/repositories/users/models"
)

const get = `
SELECT id, username, is_admin FROM users
WHERE id = ? OR username = ?;
`

type GetParams interface {
	GetId() int64
	GetUsername() string
}

func (r *UsersRepository) Get(ctx context.Context, params GetParams) (models.User, error) {

	row := r.db.QueryRowContext(ctx, get, params.GetId(), params.GetUsername())

	var user models.User
	var isAdmin int64

	err := row.Scan(&user.Id, &user.Username, &isAdmin)
	if err != nil {
		return user, err
	}

	if isAdmin == 1 {
		user.IsAdmin = true
	}

	return user, nil
}
