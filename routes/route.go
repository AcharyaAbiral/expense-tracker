package routes

import (
	"expense_tracker/handler"

	"github.com/labstack/echo/v5"
)

func RegisterExpenseRoutes(e *echo.Group, h *handler.ExpenseHandler) {
	e.POST("/expense", h.Create)
	e.GET("/expense/:id", h.FindByID)
	e.GET("/expense", h.GetExpenses)
	e.PUT("/expense/:id", h.UpdateExpense)
	e.DELETE("/expense/:id", h.DeleteExpense)
}

func RegisterUserRoutes(e *echo.Group, h *handler.UserHandler) {
	e.GET("/user/:id", h.GetUser)
}

func RegisterAuthRoutes(e *echo.Echo, h *handler.AuthHandler) {
	e.POST("/login", h.Login)
	e.POST("/signup", h.Signup)
}

func RegisterCategoryRoutes(e *echo.Echo, h *handler.CategoryHandler) {
	e.GET("/category/:id", h.FindByID)
	e.POST("/category", h.Create)
}
