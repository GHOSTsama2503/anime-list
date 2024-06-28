package sessions

import (
	"context"
	"database/sql"
	"errors"
)

const deleteQuery = `
DELETE FROM sessions
WHERE user_id = ? AND session_id = ?;
`

type DeleteSessionI interface {
	GetUserId() int64
	GetSessionId() string
}

func (repository *SessionsRepository) Delete(ctx context.Context, params DeleteSessionI) (bool, error) {

	var deleted bool

	_, err := repository.db.ExecContext(ctx, deleteQuery, params.GetUserId(), params.GetSessionId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return deleted, nil
		}

		return deleted, err
	}

	deleted = true

	return deleted, nil
}
