package animes

import (
	"context"
	"database/sql"

	"github.com/ghostsama2503/anime-list/repositories/animes/models"
)

const getList = `
SELECT id, title_romaji, description FROM animes
WHERE (
	@title = NULL
	OR title_romaji LIKE '%' || @title || '%'
	OR title_native LIKE '%' || @title || '%'
	OR title_english LIKE '%' || @title || '%'
) OR (
	@description = NULL
	OR description LIKE '%' || @description || '%'
)
ORDER BY title_romaji COLLATE NOCASE ASC
LIMIT ?
OFFSET ?;
`

type GetListParams interface {
	GetQuery() string
	GetLimit() int64
	GetOffset() int64
	GetIncludeDescriptions() bool
}

func (r *AnimesRepository) GetList(ctx context.Context, params GetListParams) ([]models.AnimeTiny, error) {

	description := sql.NullString{}
	if params.GetIncludeDescriptions() {
		description.String = params.GetQuery()
		description.Valid = params.GetIncludeDescriptions()
	}

	rows, err := r.db.QueryContext(ctx, getList,
		sql.Named("title", params.GetQuery()),
		sql.Named("description", description),
		params.GetLimit(),
		params.GetOffset(),
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
