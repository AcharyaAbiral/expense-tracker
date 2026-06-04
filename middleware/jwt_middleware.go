package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type JwtCustomClaims struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}

func JwtMiddleware(signingKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(signingKey),
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
	})
}
