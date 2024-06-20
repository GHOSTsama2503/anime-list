package genres

import (
	"anime-list/repositories/genres/models"
	"context"
)

const search = `
SELECT id, name FROM genres
WHERE name LIKE '%' || ? || '%';
`

func (r *GenresRepository) Search(ctx context.Context, name string) (models.Genre, error) {

	row := r.db.QueryRowContext(ctx, search, name)

	var genre models.Genre

	err := row.Scan(&genre.Id, &genre.Name)

	return genre, err
}
