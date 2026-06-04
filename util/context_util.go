package util

import (
	"expense_tracker/middlewares"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func GetUserID(c *echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*middlewares.JwtCustomClaims)
	return claims.UserID
}
