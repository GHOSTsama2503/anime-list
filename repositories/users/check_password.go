package users

import (
	"context"
	"database/sql"
	"errors"
)

const checkUserPassword = `
SELECT id FROM users
WHERE username = ? AND password = ?;
`

type CheckPasswordParams interface {
	GetUsername() string
	GetPassword() string
}

func (r *UsersRepository) CheckPassword(ctx context.Context, params CheckPasswordParams) (bool, error) {

	row := r.db.QueryRowContext(ctx, checkUserPassword, params.GetUsername(), params.GetPassword())

	var id int64

	err := row.Scan(&id)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return id != 0, err
}
