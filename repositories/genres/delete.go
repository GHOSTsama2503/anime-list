package genres

import "context"

const delete = `
DELETE FROM genres
WHERE id = ?;
`

func (r *GenresRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, delete, id)
	return err
}
