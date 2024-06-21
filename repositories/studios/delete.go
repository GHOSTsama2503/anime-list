package studios

import "context"

const delete = `
DELETE FROM studios
WHERE id = ?;
`

func (r *StudiosRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, delete, id)
	return err
}
