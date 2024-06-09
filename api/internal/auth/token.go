package auth

import (
	"anime-list/internal/env"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(username string) (string, error) {

	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      now.Unix(),
		"exp":      now.AddDate(0, 1, 0).Unix(),
	})

	return token.SignedString(env.JWTSecret)
}

func IsValidToken(tokenStr string) bool {

	_, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return env.JWTSecret, nil
	})

	return err == nil
}
