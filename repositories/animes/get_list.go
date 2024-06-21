package animes

import (
	"github.com/ghostsama2503/anime-list/api/repositories/animes/models"
	"context"
)

const getList = `
SELECT id, title_romaji, description FROM animes
WHERE (
	@title = NULL
	OR title_romaji LIKE '%' || @title || '%'
	OR title_native LIKE '%' || @title || '%'
	OR title_english LIKE '%' || @title || '%'
)
	OR (
	@description = NULL
	OR description LIKE '%' || @description || '%'
)
ORDER BY title_romaji COLLATE NOCASE ASC
LIMIT ?
OFFSET ?;
`

type GetListParams struct {
	Query               string
	IncludeDescriptions bool
	Limit               int64
	Offset              int64
}

func (r *AnimesRepository) GetList(ctx context.Context, params GetListParams) ([]models.AnimeTiny, error) {

	var description string
	if params.IncludeDescriptions {
		description = params.Query
	}

	rows, err := r.db.QueryContext(ctx, getList,
		params.Query,
		description,
		params.Limit,
		params.Offset,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.AnimeTiny

	for rows.Next() {
		i := models.AnimeTiny{}

		if err := rows.Scan(
			&i.Id,
			&i.Title,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
