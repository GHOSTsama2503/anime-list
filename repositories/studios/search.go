package studios

import (
	"github.com/ghostsama2503/anime-list/api/repositories/studios/models"
	"context"
)

const search = `
SELECT id, name FROM studios
WHERE name = ?;
`

func (r *StudiosRepository) Search(ctx context.Context, name string) (models.Studio, error) {

	row := r.db.QueryRowContext(ctx, search, name)

	var studio models.Studio

	err := row.Scan(&studio.Id, &studio.Name)

	return studio, err
}
