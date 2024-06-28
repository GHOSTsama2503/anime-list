package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/services/sessions"
)

type SignOutRequest struct {
	Session http.Cookie `cookie:"session"`
}

type SignOutResponse struct {
	Body struct {
		Ok bool `json:"ok"`
	}
}

func (controller *AuthController) SignOut(ctx context.Context, request *SignOutRequest) (*SignOutResponse, error) {

	sessionCookie, err := controller.sessionsService.ParseSessionCookie(request.Session.Value)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid session")
	}

	deleteSessionParams := sessions.DeleteSessionParams{
		SessionId: sessionCookie.SessionId,
		UserId:    sessionCookie.UserId,
	}

	deleted, err := controller.sessionsService.Delete(ctx, deleteSessionParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, huma.Error404NotFound("session not found")
		}

		return nil, httperr.New(http.StatusInternalServerError, "error deleting session", err)
	}

	response := &SignOutResponse{}
	response.Body.Ok = deleted

	return response, nil
}
