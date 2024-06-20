package covers

import "context"

const delete = `
DELETE FROM covers
WHERE id = ?;
`

func (r *CoversRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, delete, id)
	return err
}
