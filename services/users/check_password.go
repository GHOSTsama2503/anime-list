package users

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type CheckPasswordParams struct {
	Id       int64
	Password string
}

func (service *UsersService) CheckPassword(ctx context.Context, params CheckPasswordParams) (bool, error) {

	hashedPassword, err := service.repository.GetPassword(ctx, params.Id)
	if err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(params.Password)); err != nil {
		return false, err
	}

	return true, nil
}
