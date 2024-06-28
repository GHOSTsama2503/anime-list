package sessions

import (
	"context"

	"github.com/ghostsama2503/anime-list/repositories/sessions/models"
)

const createQuery = `
INSERT INTO sessions (user_id, session_id, session_name)
VALUES (?, ?, ?)
RETURNING id, user_id, session_id, session_name, created_at;
`

type CreateSessionI interface {
	GetSessionId() string
	GetSessionName() string
	GetUserId() int64
}

func (repository *SessionsRepository) Create(ctx context.Context, params CreateSessionI) (SessionModelI, error) {

	row := repository.db.QueryRowContext(ctx, createQuery,
		params.GetUserId(),
		params.GetSessionId(),
		params.GetSessionName(),
	)

	var session models.Session

	err := row.Scan(&session.Id, &session.UserId, &session.SessionId, &session.SessionName, &session.CreatedAt)

	return session, err
}
