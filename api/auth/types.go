package auth

import (
	"anime-list/database/queries"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (p *CreateUserParams) New(user AuthRequest, isAdmin bool) {
	p.Username = user.Body.Username
	p.Password = user.Body.Password
	p.IsAdmin = isAdmin
}

func (p *CreateUserParams) GetModel() (user queries.CreateUserParams) {
	user.Username = p.Username
	user.Password = p.Password

	if p.IsAdmin {
		user.IsAdmin = 1
	}

	return
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u *User) New(user queries.User) {
	u.Id = user.ID
	u.Username = user.Username

	if user.IsAdmin == 1 {
		u.IsAdmin = true
	}
}
