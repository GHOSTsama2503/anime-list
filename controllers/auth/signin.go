package auth

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/ghostsama2503/anime-list/common/validators"
	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/services/sessions"
	"github.com/ghostsama2503/anime-list/services/users"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	UserAgent string `header:"User-Agent"`
	Body      struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

func (request *SignInRequest) GetId() int64 {
	return 0
}

func (request *SignInRequest) GetUsername() string {
	return request.Body.Username
}

func (request *SignInRequest) GetPassword() string {
	return request.Body.Password
}

type SignInResponse struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
	Body      struct {
		Ok bool `json:"bool"`
	}
}

func (c *AuthController) SignIn(ctx context.Context, request *SignInRequest) (*SignInResponse, error) {

	if validators.StringIsEmpty(request.Body.Username) || validators.StringIsEmpty(request.Body.Password) {
		return nil, EmptyUsernameOrPasswordError
	}

	user, err := c.usersRepository.Get(ctx, request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, InvalidUsernameOrPasswordError
		}

		return nil, httperr.New(http.StatusInternalServerError, "error getting user info", err)
	}

	checkPasswordParams := users.CheckPasswordParams{
		Id:       user.Id,
		Password: request.Body.Password,
	}

	isValidPassword, err := c.usersService.CheckPassword(ctx, checkPasswordParams)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, InvalidUsernameOrPasswordError
		}

		return nil, httperr.New(http.StatusInternalServerError, "errro checking user password", err)
	}

	if !isValidPassword {
		return nil, InvalidUsernameOrPasswordError
	}

	sessionName, err := GetSessionName(request.UserAgent)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error getting session name", err)
	}

	sessionParams := sessions.CreateSessionParams{
		SessionName: sessionName,
		UserId:      user.Id,
	}

	session, err := c.sessionsService.CreateSession(ctx, sessionParams)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error creating user session", err)
	}

	cookieParams := sessions.SessionCookie{
		SessionId: session.SessionId,
		UserId:    session.UserId,
	}

	sessionId, err := c.sessionsService.NewSessionCookie(cookieParams)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error generating session id", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: sessionId,
	}

	resp := &SignInResponse{}
	resp.SetCookie = sessionCookie
	resp.Body.Ok = true

	return resp, nil
}
