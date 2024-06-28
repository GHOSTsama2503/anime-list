package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/httperr"
	"github.com/ghostsama2503/anime-list/repositories/users"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

type SignUpResponse struct {
	Body struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		IsAdmin  bool   `json:"isAdmin"`
	}
}

func (controller *AuthController) SignUp(ctx context.Context, request *SignUpRequest) (*SignUpResponse, error) {

	adminExists, err := controller.usersRepository.CheckIfAdminExists(ctx)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error checking if admin exists", err)
	}

	if adminExists {
		return nil, huma.Error401Unauthorized("only admins can do this")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Body.Password), bcrypt.DefaultCost)
	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return nil, huma.Error400BadRequest("password too long")
		}

		return nil, httperr.New(http.StatusInternalServerError, "error generating password hash", err)
	}

	createUserParams := users.CreateParams{
		Username: request.Body.Username,
		Password: string(password),
		IsAdmin:  true,
	}

	user, err := controller.usersRepository.Create(ctx, createUserParams)
	if err != nil {
		return nil, httperr.New(http.StatusInternalServerError, "error creating user", err)
	}

	response := &SignUpResponse{}
	response.Body.Id = user.Id
	response.Body.Username = user.Username
	response.Body.IsAdmin = user.IsAdmin

	return response, nil
}
