package genres

import (
	"github.com/ghostsama2503/anime-list/repositories/genres/models"
	"context"
)

const create = `
INSERT INTO genres (name)
VALUES (?)
RETURNING id, name;
`

func (r *GenresRepository) Create(ctx context.Context, name string) (models.Genre, error) {

	row := r.db.QueryRowContext(ctx, create, name)

	var genre models.Genre

	err := row.Scan(&genre.Id, &genre.Name)

	return genre, err
}
