package auth_test

import (
	"anime-list/internal/auth"
	"fmt"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestSignInHandler(t *testing.T) {
	_, api := humatest.New(t)

	api.UseMiddleware(auth.NewAuthMiddleware(api))
	auth.Use(api)

	res := api.Get("/auth")
	if res.Code != http.StatusUnauthorized {
		t.Fatalf("auth does not work, response with status: %s", http.StatusText(res.Code))
	}

	creds := auth.UserAuth{
		Username: "test",
		Password: "test",
	}

	signRes := api.Post("/signin", creds)
	if signRes.Code != http.StatusOK {
		t.Fatalf("signin response with status: %s", http.StatusText(signRes.Code))
	}

	var token string

	cookies := signRes.Result().Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
			break
		}
	}

	if token == "" {
		t.Fatal("token not found in signin cookies")
	}

	cookieHeader := fmt.Sprintf("Cookie: token=%s", token)

	authRes := api.Get("/auth", cookieHeader)
	if authRes.Code != http.StatusOK {
		t.Fatalf("auth response with status: %s", http.StatusText(authRes.Code))
	}

	fakeCookieHeader := fmt.Sprintf("Cookie: token=%s", "super reliable user")

	fakeAuthRes := api.Get("/auth", fakeCookieHeader)
	if fakeAuthRes.Code != http.StatusUnauthorized {
		t.Fatalf("token validation does not work, response with status: %s", http.StatusText(fakeAuthRes.Code))
	}
}
