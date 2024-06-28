package users

import (
	"context"
	"database/sql"
	"errors"
)

const getPasswordQuery = `
SELECT password FROM users
WHERE id = ?;
`

func (repository *UsersRepository) GetPassword(ctx context.Context, id int64) (string, error) {

	row := repository.db.QueryRowContext(ctx, getPasswordQuery, id)

	var password string

	err := row.Scan(&password)

	if errors.Is(err, sql.ErrNoRows) {
		return password, nil
	}

	return password, err
}
