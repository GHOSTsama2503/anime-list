package covers

import (
	"anime-list/repositories/covers/models"
	"context"
)

const get = `
SELECT extra_large, large, medium, color FROM cover_images
WHERE id = ?;
`

func (r *CoversRepository) Get(ctx context.Context, id int64) (models.CoverImage, error) {

	row := r.db.QueryRowContext(ctx, get, id)

	var coverImage models.CoverImage

	err := row.Scan(
		&coverImage.Id,
		&coverImage.ExtraLarge,
		&coverImage.Large,
		&coverImage.Medium,
	)

	return coverImage, err
}
