package sessions

import (
	"context"

	"github.com/ghostsama2503/anime-list/repositories/sessions"
)

func (service *SessionsService) RecoverSession(ctx context.Context, sessionId string) (sessions.SessionModelI, error) {
	return service.repository.GetBySessionId(ctx, sessionId)
}
