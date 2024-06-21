package genres

import (
	"github.com/ghostsama2503/anime-list/repositories/genres/models"
	"context"
)

const update = `
UPDATE genres
SET name = ?
WHERE id = ?
RETURNING id, name;
`

type UpdateParams struct {
	Id   int64
	Name string
}

func (r *GenresRepository) Update(ctx context.Context, params UpdateParams) (models.Genre, error) {

	row := r.db.QueryRowContext(ctx, update, params.Name, params.Id)

	var genre models.Genre

	err := row.Scan(&genre.Id, &genre.Name)

	return genre, err
}
