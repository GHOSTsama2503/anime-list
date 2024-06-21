package users

import "context"

const checkUserPassword = `
SELECT id FROM users
WHERE username = ? AND password = ?;
`

type CheckUserPasswordParams struct {
	Username string
	Password string
}

func (r *UsersRepository) CheckPassword(ctx context.Context, params CheckUserPasswordParams) bool {

	row := r.db.QueryRowContext(ctx, checkUserPassword, params.Username, params.Password)

	var id int64

	row.Scan(&id)

	return id != 0
}
