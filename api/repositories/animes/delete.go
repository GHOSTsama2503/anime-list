package animes

import "context"

const delete = `
DELETE FROM animes
WHERE id = ?;
`

func (r *AnimesRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, delete, id)
	return err
}
