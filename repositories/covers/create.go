package covers

import (
	"github.com/ghostsama2503/anime-list/repositories/covers/models"
	"context"
)

const create = `
INSERT INTO cover_images (extra_large, large, medium, color)
VALUES (?, ?, ?, ?)
RETURNING id, extra_large, large, medium, color;
`

type CreateParams struct {
	ExtraLarge string
	Large      string
	Medium     string
	Color      string
}

func (r *CoversRepository) Create(ctx context.Context, params CreateParams) (models.CoverImage, error) {

	row := r.db.QueryRowContext(ctx, create,
		params.ExtraLarge,
		params.Large,
		params.Medium,
		params.Color,
	)

	var coverImage models.CoverImage

	err := row.Scan(
		&coverImage.Id,
		&coverImage.ExtraLarge,
		&coverImage.Large,
		&coverImage.Medium,
		&coverImage.Color,
	)

	return coverImage, err
}
