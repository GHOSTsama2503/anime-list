package genres

import (
	"github.com/ghostsama2503/anime-list/api/repositories/genres/models"
	"context"
)

const get = `
SELECT id, name FROM genres
WHERE id = ?;
`

func (r *GenresRepository) Get(ctx context.Context, id int64) (models.Genre, error) {

	row := r.db.QueryRowContext(ctx, get, id)

	var genre models.Genre

	err := row.Scan(&genre.Id, &genre.Name)

	return genre, err
}
