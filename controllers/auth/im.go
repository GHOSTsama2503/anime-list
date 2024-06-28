package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/httperr"
)

type ImRequest struct {
	SessionId http.Cookie `cookie:"sessionId"`
}

type ImResponse struct {
	Body struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		IsAdmin  bool   `json:"isAdmin"`
	}
}

type GetUserParams struct {
	Id int64
}

func (params GetUserParams) GetId() int64 {
	return params.Id
}

func (params GetUserParams) GetUsername() string {
	return ""
}

func (controller *AuthController) Im(ctx context.Context, request *ImRequest) (*ImResponse, error) {

	session, err := controller.sessionsService.RecoverSession(ctx, request.SessionId.Value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, huma.NewError(http.StatusTeapot, http.StatusText(http.StatusTeapot))
		}

		return nil, httperr.New(http.StatusInternalServerError, "error recovering session", err)
	}

	getUserParams := GetUserParams{
		Id: session.GetUserId(),
	}

	user, err := controller.usersRepository.Get(ctx, getUserParams)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting user info", err)
	}

	response := &ImResponse{}
	response.Body.Id = user.Id
	response.Body.Username = user.Username
	response.Body.IsAdmin = user.IsAdmin

	return response, nil
}
