package users

import (
	"context"
	"database/sql"
	"errors"
)

const checkIfAdminExists = `
SELECT id FROM users
WHERE is_admin = 1;
`

func (r *UsersRepository) CheckIfAdminExists(ctx context.Context) (bool, error) {

	row := r.db.QueryRowContext(ctx, checkIfAdminExists)

	var id int64

	err := row.Scan(&id)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return id != 0, err
}
