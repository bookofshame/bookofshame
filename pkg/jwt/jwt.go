package jwt

import (
	"context"
	"fmt"
	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

type Jwt struct {
	secret string
	logger *zap.SugaredLogger
}

func New(ctx context.Context, cfg config.Config) *Jwt {
	return &Jwt{
		secret: cfg.JwtSecret,
		logger: logging.FromContext(ctx),
	}
}

func (j Jwt) Token(pl Payload) (string, error) {
	claims := Claims{
		pl.UserId,
		pl.UserLocale,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secret))

	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token. error: %w", err)
	}

	return signedToken, nil
}

func (j Jwt) Parse(tokenStr string) (*Payload, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("failed to validate signing method")
		}

		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed parsing jwt. error: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	pl := claims.Payload()

	return pl, nil
}

func GetDataFromContext(ctx context.Context) (Payload, error) {
	data := ctx.Value(dataCtxKey)
	pl, ok := data.(*Payload)
	if !ok {
		return Payload{}, fmt.Errorf("jwt data not found in context")
	}

	return *pl, nil
}
