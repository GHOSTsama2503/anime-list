package sessions

import "context"

type ValidateParams interface {
	GetSessionId() string
	GetUserId() int64
}

func (service *SessionsService) Validate(ctx context.Context, sessionId string) (bool, error) {

	var isValidSession bool

	_, err := service.repository.GetBySessionId(ctx, sessionId)
	if err != nil {
		return isValidSession, err
	}

	return !isValidSession, nil
}
