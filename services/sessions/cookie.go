package sessions

import (
	"encoding/json"
	"time"

	"github.com/ghostsama2503/anime-list/common/config"
	"github.com/ghostsama2503/go-jwt"
)

type SessionCookie struct {
	SessionId string
	UserId    int64
}

type SessionClaims struct {
	Issuer         string `json:"iss"`
	Subject        string `json:"sub"`
	ExpirationTime int64  `json:"exp"`
	IssuedAtTime   int64  `json:"iat"`
	SessionId      string `json:"sessionId"`
	UserId         int64  `json:"userId"`
}

func (claims *SessionClaims) Marshal() ([]byte, error) {
	return json.Marshal(claims)
}

func (claims *SessionClaims) GetIssuer() string {
	return claims.Issuer
}

func (claims *SessionClaims) GetSubject() string {
	return claims.Subject
}

func (claims *SessionClaims) GetExpirationTime() int64 {
	return claims.ExpirationTime
}

func (claims *SessionClaims) GetIssuedAtTime() int64 {
	return claims.IssuedAtTime
}

func (sessions *SessionsService) NewSessionCookie(params SessionCookie) (string, error) {

	var sessionCookie string

	currentTime := time.Now().UTC()

	claims := &SessionClaims{}
	claims.SessionId = params.SessionId
	claims.UserId = params.UserId
	claims.IssuedAtTime = currentTime.Unix()
	claims.ExpirationTime = currentTime.AddDate(0, 1, 0).Unix()

	token := jwt.New(claims)

	signedToken, err := token.Signed([]byte(config.Env.JWTSecret))
	if err != nil {
		return sessionCookie, err
	}

	return signedToken, nil
}

func (sessions *SessionsService) ParseSessionCookie(cookie string) (SessionCookie, error) {

	var sessionCookie SessionCookie

	claims := &SessionClaims{}

	token, err := jwt.ParseWithClaims(cookie, claims)
	if err != nil {
		return sessionCookie, err
	}

	if err := token.Validate([]byte(config.Env.JWTSecret)); err != nil {
		return sessionCookie, err
	}

	sessionCookie.SessionId = claims.SessionId
	sessionCookie.UserId = claims.UserId

	return sessionCookie, err
}
