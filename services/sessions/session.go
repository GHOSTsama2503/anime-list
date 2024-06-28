package sessions

import (
	"time"

	"github.com/ghostsama2503/anime-list/repositories/sessions"
)

type Session struct {
	Id          int64
	SessionId   string
	SessionName string
	UserId      int64
	CreatedAt   time.Time
}

func NewFromModel(model sessions.SessionModelI) Session {
	return Session{
		Id:          model.GetId(),
		SessionId:   model.GetSessionId(),
		SessionName: model.GetSessionName(),
		UserId:      model.GetUserId(),
		CreatedAt:   model.GetCreatedAt(),
	}
}
