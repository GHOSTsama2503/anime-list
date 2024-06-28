package models

import "time"

type Session struct {
	Id          int64
	SessionId   string
	SessionName string
	UserId      int64
	CreatedAt   time.Time
}

func (session Session) GetId() int64 {
	return session.Id
}

func (session Session) GetSessionId() string {
	return session.SessionId
}

func (session Session) GetSessionName() string {
	return session.SessionName
}

func (session Session) GetUserId() int64 {
	return session.UserId
}

func (session Session) GetCreatedAt() time.Time {
	return session.CreatedAt
}
