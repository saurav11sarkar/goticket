package event

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/saurav11sarkar/ticket/internal/event/dto"
	"github.com/saurav11sarkar/ticket/internal/httpResponse"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *echo.Context) error {
	var req dto.CreateEventRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Detail:  err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Detail:  err.Error(),
		})
	}

	data, err := h.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Service error",
			Detail:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, httpResponse.Success{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    data,
	})
}
