package user

import (
	"github.com/labstack/echo/v5"
	"github.com/saurav11sarkar/ticket/internal/auth"
	"gorm.io/gorm"
)

func RegisterRouter(e *echo.Echo, db *gorm.DB) {
	userRepository := NewUserRepository(db)
	jwtService := auth.NewJwtService("sdfjisgh", 15)
	userService := NewUserService(userRepository, jwtService)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
}
