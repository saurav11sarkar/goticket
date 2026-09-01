package user

import (
	"time"

	"github.com/labstack/echo/v5"
	"github.com/saurav11sarkar/ticket/internal/auth"
	"github.com/saurav11sarkar/ticket/internal/config"
	"github.com/saurav11sarkar/ticket/internal/middlewares"
	"github.com/saurav11sarkar/ticket/internal/utils"
	"gorm.io/gorm"
)

func RegisterRouter(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	userRepository := NewUserRepository(db)
	jwtService := auth.NewJwtService("sdfjisgh", 15*time.Minute)
	emailSender := utils.NewEmailSender(cfg)
	userService := NewUserService(userRepository, jwtService, emailSender)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService))
}
