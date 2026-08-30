package event

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRouter(e *echo.Echo, db *gorm.DB) {
	eventRepository := NewRepository(db)
	eventService := NewService(eventRepository)
	eventHandler := NewHandler(eventService)

	api := e.Group("/api/v1/event")

	api.POST("/", eventHandler.Create)
	api.GET("/", eventHandler.GetAll)
	api.GET("/:id", eventHandler.GetByID)
	api.PUT("/:id", eventHandler.Update)
	api.PATCH("/:id", eventHandler.Update)
	api.DELETE("/:id", eventHandler.Delete)

}
