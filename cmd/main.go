package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	return nil
}

type CustomValidator struct {
	validate *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validate.Struct(i)
}

type SuccessResponse struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Errors     any    `json:"errors,omitempty"`
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validate: validator.New()}

	dsn := "host=localhost user=postgres password=123456 dbname=tricket port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		e.Logger.Error("failed to connect database", "error", err)
		return
	}
	db.AutoMigrate(&User{})
	e.Logger.Info("Connected to database")

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, SuccessResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Message:    "Wellcome to Go server!",
			Data:       nil,
		})
	})

	e.POST("/", func(c *echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Success:    false,
				Message:    "Invalid request",
				Errors:     err.Error(),
			})
		}

		if err := c.Validate(user); err != nil {
			errorMap := make(map[string]string)
			for _, vErr := range err.(validator.ValidationErrors) {
				errorMap[vErr.Field()] = vErr.Error()
			}
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Success:    false,
				Message:    "Validation Failed",
				Errors:     errorMap,
			})
		}

		if result := db.Create(&user); result.Error != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Success:    false,
				Message:    "Failed to create user",
				Errors:     result.Error.Error(),
			})
		}

		return c.JSON(http.StatusOK, SuccessResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Message:    "User Create successfully",
			Data:       &user,
		})
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
