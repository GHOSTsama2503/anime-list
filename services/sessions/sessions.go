package sessions

import (
	"context"

	"github.com/ghostsama2503/anime-list/repositories/sessions"
)

type SessionsRepository interface {
	Create(ctx context.Context, params sessions.CreateSessionI) (sessions.SessionModelI, error)
	GetBySessionId(ctx context.Context, sessionId string) (sessions.SessionModelI, error)
	Delete(ctx context.Context, params sessions.DeleteSessionI) (bool, error)
}

type SessionsService struct {
	repository SessionsRepository
}

func New(repository SessionsRepository) *SessionsService {
	return &SessionsService{repository}
}
