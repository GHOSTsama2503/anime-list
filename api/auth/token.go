package auth

import (
	"anime-list/common/config"
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

	return token.SignedString([]byte(config.Env.JWTSecret))
}

func IsValidToken(tokenStr string) bool {

	_, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Env.JWTSecret), nil
	})

	return err == nil
}
