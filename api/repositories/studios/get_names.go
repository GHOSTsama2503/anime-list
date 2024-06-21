package studios

import "context"

const getNames = `
SELECT name FROM studios
WHERE id IN (
	SELECT studio_id FROM anime_studios
	WHERE anime_id = ?
);
`

func (r *StudiosRepository) GetNames(ctx context.Context, animeId int64) ([]string, error) {

	var studios []string

	rows, err := r.db.QueryContext(ctx, getNames, animeId)
	if err != nil {
		return studios, err
	}

	for rows.Next() {
		var studio string
		if err := rows.Scan(&studio); err != nil {
			return studios, err
		}

		studios = append(studios, studio)
	}

	return studios, err
}
