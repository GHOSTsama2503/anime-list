package auth

import "net/http"

type AuthRequest struct {
	Body UserAuth
}

type SignInResponse struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
	Body      struct {
		Ok bool `json:"ok"`
	}
}

type SignUpResponse struct {
	Body User `json:"user"`
}

type TestAuthRequest struct {
}

type TestAuthResponse struct {
	Body struct {
		Ok bool `json:"ok"`
	}
}
