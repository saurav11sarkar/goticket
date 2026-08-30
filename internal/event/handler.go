package event

import (
	"errors"
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

func (h *Handler) GetAll(c *echo.Context) error {
	data, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Service error",
			Detail:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, httpResponse.Success{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func (h *Handler) GetByID(c *echo.Context) error {
	data, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return c.JSON(http.StatusNotFound, httpResponse.Error{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get event",
			Detail:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, httpResponse.Success{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func (h *Handler) Update(c *echo.Context) error {
	var req dto.UpdateEventRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Detail:  err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Detail:  err.Error(),
		})
	}

	data, err := h.service.Update(c.Param("id"), req)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return c.JSON(http.StatusNotFound, httpResponse.Error{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update event",
			Detail:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpResponse.Success{
		Code:    http.StatusOK,
		Message: "Event updated",
		Data:    data,
	})
}

func (h *Handler) Delete(c *echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		if errors.Is(err, ErrNotFound) {
			return c.JSON(http.StatusNotFound, httpResponse.Error{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete event",
			Detail:  err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
