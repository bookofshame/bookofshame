package jwt

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	UserId     int
	UserLocale string
}

type Claims struct {
	UserId     int    `json:"userId"`
	UserLocale string `json:"userLocale"`
	jwt.RegisteredClaims
}

func (c Claims) Payload() *Payload {
	return &Payload{
		UserId:     c.UserId,
		UserLocale: c.UserLocale,
	}
}

type ContextKey string

const (
	dataCtxKey ContextKey = "jwtData"
)
