package sessions

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type CreateSessionParams struct {
	SessionName string
	UserId      int64
}

type SessionParams struct {
	SessionId   string
	SessionName string
	UserId      int64
	CreatedAt   int64
}

func (session SessionParams) GetSessionId() string {
	return session.SessionId
}

func (session SessionParams) GetSessionName() string {
	return session.SessionName
}

func (session SessionParams) GetUserId() int64 {
	return session.UserId
}

func (session SessionParams) GetCreatedAt() int64 {
	return session.CreatedAt
}

func (service *SessionsService) CreateSession(ctx context.Context, params CreateSessionParams) (Session, error) {

	var session Session

	sessionId, err := uuid.NewRandom()
	if err != nil {
		return session, fmt.Errorf("creating session id: %v", err)
	}

	sessionParams := SessionParams{
		UserId:      params.UserId,
		SessionId:   sessionId.String(),
		SessionName: params.SessionName,
	}

	sessionModel, err := service.repository.Create(ctx, sessionParams)
	if err != nil {
		return session, err
	}

	session = NewFromModel(sessionModel)

	return session, nil
}
