package covers

import (
	"github.com/ghostsama2503/anime-list/api/repositories/covers/models"
	"context"
)

const get = `
SELECT extra_large, large, medium, color FROM cover_images
WHERE id = ?;
`

func (r *CoversRepository) Get(ctx context.Context, id int64) (models.CoverImage, error) {

	var coverImage models.CoverImage

	rows, err := r.db.QueryContext(ctx, get, id)
	if err != nil {
		return coverImage, err
	}

	for rows.Next() {
		err := rows.Scan(
			&coverImage.Id,
			&coverImage.ExtraLarge,
			&coverImage.Large,
			&coverImage.Medium,
		)

		if err != nil {
			return coverImage, err
		}
	}

	return coverImage, err
}
