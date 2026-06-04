package handler

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/service"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Login(c *echo.Context) error {
	var reqDto dto.LoginRequest

	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	authResponse, err := h.service.Login(&reqDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, authResponse)

}

func (h *AuthHandler) Signup(c *echo.Context) error {
	var reqDto dto.SignupRequest

	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.Signup(&reqDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, mapper.ToUserResponse(user))
}
