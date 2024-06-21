package covers

import (
	"github.com/ghostsama2503/anime-list/api/repositories/covers/models"
	"context"
)

const getData = `
SELECT extra_large, large, medium, color FROM cover_images
WHERE anime_id = ?;
`

func (r *CoversRepository) GetData(ctx context.Context, animeId int64) (models.CoverData, error) {

	var data models.CoverData

	row, err := r.db.QueryContext(ctx, getData, animeId)
	if err != nil {
		return data, err
	}

	for row.Next() {
		err := row.Scan(&data.ExtraLarge, &data.Large, &data.Medium, &data.Color)
		if err != nil {
			return data, err
		}
	}

	return data, err
}
