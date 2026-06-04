package handler

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/service"
	"expense_tracker/util"
	"net/http"
	"strconv"
	"time"

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
func (h *ExpenseHandler) GetExpenses(c *echo.Context) error {
	var paginationInput dto.PaginationInput

	if err := c.Bind(&paginationInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var categoryID *uint
	categoryIdStr := c.QueryParam("category_id")

	if categoryIdStr != "" {
		id, err := strconv.ParseUint(categoryIdStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid category_id")
		}
		tmp := uint(id)
		categoryID = &tmp
	}

	userID := util.GetUserID(c)

	expenses, err := h.service.GetExpenses(categoryID, userID, paginationInput)

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

func (h *ExpenseHandler) Summary(c *echo.Context) error {
	fromDateStr := c.QueryParam("from")
	toDateStr := c.QueryParam("to")

	if fromDateStr == "" || toDateStr == "" {
		return c.JSON(http.StatusBadRequest, "both from and to required")
	}

	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format")
	}

	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format")
	}

	userID := util.GetUserID(c)

	summary, err := h.service.GetSummary(userID, fromDate, toDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, summary)
}

func (h *ExpenseHandler) YearlySummary(c *echo.Context) error {
	yearStr := c.Param("year")
	year64, err := strconv.ParseUint(yearStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid year")
	}

	userID := util.GetUserID(c)

	summary, err := h.service.GetYearlySummary(userID, int(year64))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, summary)
}
