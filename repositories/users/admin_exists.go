package users

import "context"

const checkIfAdminExists = `
SELECT id FROM users
WHERE is_admin = 1;
`

func (r *UsersRepository) CheckIfAdminExists(ctx context.Context) (bool, error) {

	row := r.db.QueryRowContext(ctx, checkIfAdminExists)

	var id int64

	err := row.Scan(&id)

	return id != 0, err
}
