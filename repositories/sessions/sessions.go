package sessions

import (
	"time"

	"github.com/ghostsama2503/anime-list/database"
)

type SessionModelI interface {
	GetId() int64
	GetSessionId() string
	GetSessionName() string
	GetUserId() int64
	GetCreatedAt() time.Time
}

type SessionsRepository struct {
	db database.DBTX
}

func New(db database.DBTX) *SessionsRepository {
	return &SessionsRepository{db}
}
