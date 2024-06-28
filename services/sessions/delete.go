package sessions

import "context"

type DeleteSessionParams struct {
	SessionId string
	UserId    int64
}

func (params DeleteSessionParams) GetSessionId() string {
	return params.SessionId
}

func (params DeleteSessionParams) GetUserId() int64 {
	return params.UserId
}

func (service *SessionsService) Delete(ctx context.Context, params DeleteSessionParams) (bool, error) {
	return service.repository.Delete(ctx, params)
}
