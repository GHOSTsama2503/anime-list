package studios

import (
	"github.com/ghostsama2503/anime-list/api/repositories/studios/models"
	"context"
)

const get = `
SELECT id, name FROM studios
WHERE id = ?;
`

func (r *StudiosRepository) Get(ctx context.Context, id int64) (models.Studio, error) {

	row := r.db.QueryRowContext(ctx, get, id)

	var studio models.Studio

	err := row.Scan(&studio.Id, &studio.Name)

	return studio, err
}
