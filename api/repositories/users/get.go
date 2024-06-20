package users

import (
	"anime-list/repositories/users/models"
	"context"
)

const get = `
SELECT id, username, is_admin FROM users
WHERE id = ? OR username = ?;
`

type GetParams struct {
	Id       int64
	Username string
}

func (r *UsersRepository) Get(ctx context.Context, params GetParams) (models.User, error) {

	row := r.db.QueryRowContext(ctx, get, params.Id, params.Username)

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
