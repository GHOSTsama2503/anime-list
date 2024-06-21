package animes

import (
	"context"
)

const anilist = `
SELECT id FROM animes
WHERE id_al = ?;
`

func (q *AnimesRepository) TranslateAnilistId(ctx context.Context, id int64) (int64, error) {

	var localId int64

	rows, err := q.db.QueryContext(ctx, anilist, id)
	if err != nil {
		return localId, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&localId)
		if err != nil {
			return localId, err
		}
	}

	return localId, err
}
