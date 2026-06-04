package util

import (
	"expense_tracker/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func GetUserID(c *echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*middleware.JwtCustomClaims)
	return claims.UserID
}
