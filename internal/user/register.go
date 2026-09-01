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
	cloudinaryUploader, err := utils.NewCloudinaryUploader(cfg)
	if err != nil {
		panic(err)
	}
	userService := NewUserService(userRepository, jwtService, emailSender)
	userHandler := NewHandler(userService, cloudinaryUploader)
	authMiddleware := middlewares.AuthMiddleware(jwtService)

	authRoutes := e.Group("/api/v1/auth")
	authRoutes.POST("/register", userHandler.CreateUser)
	authRoutes.POST("/login", userHandler.LoginUser)
	authRoutes.GET("/me", userHandler.GetMe, authMiddleware)

	userRoutes := e.Group("/api/v1/users", authMiddleware)
	userRoutes.GET("", userHandler.GetAll)
}
