package genres

import "context"

const getNames = `
SELECT name FROM genres
WHERE id IN (
	SELECT genre_id FROM anime_genres
	WHERE anime_id = ?
);
`

func (r *GenresRepository) GetNames(ctx context.Context, animeId int64) ([]string, error) {

	var genres []string

	rows, err := r.db.QueryContext(ctx, getNames, animeId)
	if err != nil {
		return genres, err
	}

	defer rows.Close()

	for rows.Next() {
		var genre string
		if err := rows.Scan(&genre); err != nil {
			return genres, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}
