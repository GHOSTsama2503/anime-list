package auth

import (
	"anime-list/common/config"
	"context"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
)

func SignUpHandler(ctx context.Context, input *AuthRequest) (*SignUpResponse, error) {
	res := &SignUpResponse{}

	admExists, err := AdminExists(ctx)
	if err != nil {
		log.Error(err)
		return nil, huma.Error500InternalServerError("unexpected error")
	}

	if admExists {
		return nil, huma.Error403Forbidden("")
	}

	params := CreateUserParams{}
	params.New(*input, !admExists)

	user, err := CreateUser(ctx, params)
	if err != nil {
		return nil, huma.Error409Conflict(err.Error())
	}

	res.Body = user
	return res, nil
}

func SignInHandler(ctx context.Context, input *AuthRequest) (*SignInResponse, error) {
	res := &SignInResponse{}

	user, err := CheckUser(ctx, input.Body.Username, input.Body.Password)
	if err != nil {
		return nil, huma.Error401Unauthorized(err.Error())
	}

	token, err := NewToken(user.Username)
	if err != nil {
		log.Error(err)
		return nil, huma.Error500InternalServerError("unexpected error")
	}

	res.Body.Ok = true
	res.SetCookie = http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   config.Env.IsProduction,
		Expires:  time.Now().AddDate(0, 1, 0),
	}

	return res, nil
}

func TestAuthHandler(ctx context.Context, input *TestAuthRequest) (*TestAuthResponse, error) {
	res := &TestAuthResponse{}
	res.Body.Ok = true

	return res, nil
}

func Use(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "signup",
		Method:      http.MethodPost,
		Path:        "/signup",
		Summary:     "Create new user",
		Security: []map[string][]string{
			{"auth": {"default"}},
		},
	}, SignUpHandler)

	huma.Register(api, huma.Operation{
		OperationID: "signin",
		Method:      http.MethodPost,
		Path:        "/signin",
		Summary:     "Login using credentials",
	}, SignInHandler)

	huma.Register(api, huma.Operation{
		OperationID: "auth-test",
		Method:      http.MethodGet,
		Path:        "/auth",
		Summary:     "Auth test",
		Security: []map[string][]string{
			{"auth": {"default"}},
		},
	}, TestAuthHandler)
}
