package handler

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) Create(c *echo.Context) error {
	var reqDto dto.CategoryRequest

	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	category, err := h.service.Create(&reqDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, mapper.ToCategoryResponse(category))
}

func (h *CategoryHandler) FindByID(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	category, err := h.service.FindByID(uint(id64))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, mapper.ToCategoryResponse(category))
}
