package studios

import (
	"github.com/ghostsama2503/anime-list/api/repositories/studios/models"
	"context"
)

const create = `
INSERT INTO studios (name)
VALUES (?)
RETURNING id, name;
`

func (r *StudiosRepository) Create(ctx context.Context, name string) (models.Studio, error) {

	row := r.db.QueryRowContext(ctx, create, name)

	var studio models.Studio

	err := row.Scan(&studio.Id, &studio.Name)

	return studio, err
}
