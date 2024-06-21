package users

import "context"

const delete = `
DELETE FROM users
WHERE id = ?;
`

func (r *UsersRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, delete, id)
	return err
}
