package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/saurav11sarkar/ticket/internal/httpResponse"
	"github.com/saurav11sarkar/ticket/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type customValidator struct {
	validate *validator.Validate
}

func (cv *customValidator) Validate(value any) error {
	return cv.validate.Struct(value)
}

func main() {
	e := echo.New()
	e.Validator = &customValidator{validate: validator.New()}

	dsn := "host=localhost user=postgres password=123456 dbname=tricket port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		e.Logger.Error("failed to connect database", "error", err)
		return
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		e.Logger.Error("failed to migrate database", "error", err)
		return
	}
	e.Logger.Info("Connected to database")

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpResponse.Success{
			Code:    http.StatusOK,
			Message: "Welcome to Go server!",
		})
	})

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	e.POST("/users", userHandler.CreateUser)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
