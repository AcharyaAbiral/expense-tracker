package handler

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/service"
	"expense_tracker/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(service *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service}
}

func (h *ExpenseHandler) Create(c *echo.Context) error {
	var reqDto dto.ExpenseRequest

	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID := util.GetUserID(c)

	expense, err := h.service.Create(&reqDto, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, mapper.ToExpenseResponse(expense))
}

func (h *ExpenseHandler) FindByID(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	userID := util.GetUserID(c)

	expense, err := h.service.FindByIDAndUserID(uint(id64), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, mapper.ToExpenseResponse(expense))
}

// take date range as well?? :todo
func (h *ExpenseHandler) GetExpenses(c *echo.Context) error { //take category argument as well??
	var paginationInput dto.PaginationInput

	if err := c.Bind(&paginationInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID := util.GetUserID(c)

	expenses, err := h.service.GetExpenses(userID, paginationInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) UpdateExpense(c *echo.Context) error {
	var reqDto dto.ExpenseRequest

	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userID := util.GetUserID(c)

	expense, err := h.service.Update(&reqDto, uint(id64), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, mapper.ToExpenseResponse(expense))

}

func (h *ExpenseHandler) DeleteExpense(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	userID := util.GetUserID(c)

	if err := h.service.DeleteByIDAndUserID(uint(id64), userID); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, idStr)
}
