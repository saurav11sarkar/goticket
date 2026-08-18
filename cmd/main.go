package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type CustomValidator struct {
	validate *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validate.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validate: validator.New()}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.POST("/", func(c *echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"message": "Invalid request"})
		}

		if err := c.Validate(user); err != nil {
			errorMap := make(map[string]string)
			for _, vErr := range err.(validator.ValidationErrors) {
				errorMap[vErr.Field()] = "Please check your " + vErr.Field()
			}
			return c.JSON(http.StatusBadRequest, map[string]any{
				"success": false,
				"message": "Validation Failed",
				"errors":  errorMap,
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "User Create successfully",
			"data":    user,
		})
	})

	e.Start(":8080")
}
