package main

import (
	"expense_tracker/config"
	"expense_tracker/handler"
	"expense_tracker/middleware"
	"expense_tracker/model"
	"expense_tracker/repository"
	"expense_tracker/routes"
	"expense_tracker/service"
	"expense_tracker/validator"

	"github.com/labstack/echo/v5"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	db := config.InitDB(cfg)

	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Expense{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	expRepo := repository.NewExpenseRepository(db)
	expService := service.NewExpenseService(expRepo, categoryService)
	expHandler := handler.NewExpenseHandler(expService)

	authService := service.NewAuthService(userService, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)

	g := e.Group("")
	g.Use(middleware.JwtMiddleware(cfg.JWTSecret))

	routes.RegisterExpenseRoutes(g, expHandler)
	routes.RegisterUserRoutes(g, userHandler)
	routes.RegisterCategoryRoutes(e, categoryHandler)
	routes.RegisterAuthRoutes(e, authHandler)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start", "error", err)
	}
}
