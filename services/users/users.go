package users

import "github.com/ghostsama2503/anime-list/repositories/users"

type UsersService struct {
	repository users.UsersRepositoryInterface
}

func New(repository users.UsersRepositoryInterface) *UsersService {
	return &UsersService{repository}
}
