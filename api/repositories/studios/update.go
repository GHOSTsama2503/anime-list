package studios

import (
	"anime-list/repositories/studios/models"
	"context"
)

const update = `
UPDATE studios
SET name = ?
WHERE id = ?
RETURNING id, name;
`

type UpdateParams struct {
	Id   int64
	Name string
}

func (r *StudiosRepository) Update(ctx context.Context, params UpdateParams) (models.Studio, error) {

	row := r.db.QueryRowContext(ctx, update, params.Name, params.Id)

	var studio models.Studio

	err := row.Scan(&studio.Id, &studio.Name)

	return studio, err
}
