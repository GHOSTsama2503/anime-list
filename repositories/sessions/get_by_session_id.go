package sessions

import (
	"context"

	"github.com/ghostsama2503/anime-list/repositories/sessions/models"
)

const getBySessionId = `
SELECT id, user_id, session_id, session_name, created_at FROM sessions
WHERE session_id = ?;
`

func (service *SessionsRepository) GetBySessionId(ctx context.Context, sessionId string) (SessionModelI, error) {

	row := service.db.QueryRowContext(ctx, getBySessionId, sessionId)

	var session models.Session

	err := row.Scan(&session.Id, &session.UserId, &session.SessionId, &session.SessionName, &session.CreatedAt)

	return session, err
}
