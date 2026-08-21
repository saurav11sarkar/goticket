package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRouter(e *echo.Echo, db *gorm.DB) {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewHandler(userService)

	e.POST("/users", userHandler.CreateUser)
}
