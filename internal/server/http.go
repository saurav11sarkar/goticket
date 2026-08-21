package server

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/saurav11sarkar/ticket/internal/config"
	"github.com/saurav11sarkar/ticket/internal/httpResponse"
	"github.com/saurav11sarkar/ticket/internal/user"
	"gorm.io/gorm"
)

type customValidator struct {
	validate *validator.Validate
}

func (cv *customValidator) Validate(value any) error {
	return cv.validate.Struct(value)
}

func Start(db *gorm.DB, cfg *config.Config) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	e := echo.New()
	e.Validator = &customValidator{validate: validator.New()}
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpResponse.Success{
			Code:    http.StatusOK,
			Message: "Welcome to Go server!",
		})
	})

	user.RegisterRouter(e, db)

	e.Logger.Info("Connected to database")
	return e.Start(":" + cfg.Port)
}
