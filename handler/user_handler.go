package handler

import (
	"expense_tracker/mapper"
	"expense_tracker/service"
	"expense_tracker/util"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUser(c *echo.Context) error {
	userID := util.GetUserID(c)

	user, err := h.service.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	responseDto := mapper.ToUserResponse(user)
	return c.JSON(http.StatusOK, responseDto)
}
