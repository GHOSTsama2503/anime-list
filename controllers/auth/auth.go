package auth

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/database"
	sessionsRepo "github.com/ghostsama2503/anime-list/repositories/sessions"
	usersRepo "github.com/ghostsama2503/anime-list/repositories/users"
	"github.com/ghostsama2503/anime-list/services/sessions"
	"github.com/ghostsama2503/anime-list/services/users"
)

type AuthService interface {
	NewAccessToken(subject string) (string, error)
	NewRefreshToken(subject string) (string, error)
}

type AuthController struct {
	service         AuthService
	sessionsService *sessions.SessionsService
	usersService    *users.UsersService
	usersRepository usersRepo.UsersRepositoryInterface
}

func Init(api huma.API, db database.DBTX) {

	sessionsRepository := sessionsRepo.New(db)
	sessionsService := sessions.New(sessionsRepository)

	usersRepository := usersRepo.NewUsersRepository(db)
	usersService := users.New(usersRepository)

	controller := AuthController{
		sessionsService: sessionsService,
		usersService:    usersService,
		usersRepository: usersRepository,
	}

	group := "auth"

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/auth/im",
		Summary: "get current user info",
		Tags:    []string{group},
	}, controller.Im)

	huma.Register(api, huma.Operation{
		Method:  http.MethodPost,
		Path:    "/auth/signup",
		Summary: "sign up session",
		Tags:    []string{group},
	}, controller.SignUp)

	huma.Register(api, huma.Operation{
		Method:  http.MethodPost,
		Path:    "/auth/signin",
		Summary: "sign in session",
		Tags:    []string{group},
	}, controller.SignIn)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/auth/signout",
		Summary: "sign out session",
		Tags:    []string{group},
	}, controller.SignOut)

	huma.Register(api, huma.Operation{
		Method:  http.MethodGet,
		Path:    "/auth/test",
		Summary: "test if auth is working",
		Tags:    []string{group},
		Security: []map[string][]string{
			{"cookie": {"default"}},
		},
	}, controller.Test)
}
