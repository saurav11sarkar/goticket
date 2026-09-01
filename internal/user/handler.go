package user

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/saurav11sarkar/ticket/internal/httpResponse"
	"github.com/saurav11sarkar/ticket/internal/user/dto"
	"github.com/saurav11sarkar/ticket/internal/utils"
)

type Handler struct {
	service            *Service
	cloudinaryUploader *utils.CloudinaryUploader
}

func NewHandler(service *Service, cloudinaryUploader *utils.CloudinaryUploader) *Handler {
	return &Handler{service: service, cloudinaryUploader: cloudinaryUploader}
}

func (handler *Handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequest

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
			Message: "Validation failed",
			Detail:  err.Error(),
		})
	}

	file, err := c.FormFile("profile_image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Profile image is required",
		})
	}
	profileImageURL, err := handler.cloudinaryUploader.UploadImage(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to upload image",
		})
	}
	user, err := handler.service.CreateUser(req, profileImageURL)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			return c.JSON(http.StatusConflict, httpResponse.Error{
				Code:    http.StatusConflict,
				Message: "User already exists",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
	}

	return c.JSON(http.StatusCreated, httpResponse.Success{
		Code:    http.StatusCreated,
		Message: "User created",
		Data:    user,
	})
}

func (handler *Handler) LoginUser(c *echo.Context) error {
	req := dto.LoginRequest{}
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
			Message: "Validation failed",
			Detail:  err.Error(),
		})
	}
	user, err := handler.service.LoginUser(req)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpResponse.Error{
				Code:    http.StatusUnauthorized,
				Message: "Invalid email or password",
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK, httpResponse.Success{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    user,
	})
}

func (handler *Handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("id").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, httpResponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	user, err := handler.service.GetMe(userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return c.JSON(
				http.StatusNotFound,
				httpResponse.Error{
					Code:    http.StatusNotFound,
					Message: "User not found",
				},
			)
		}

		return c.JSON(
			http.StatusInternalServerError,
			httpResponse.Error{
				Code:    http.StatusInternalServerError,
				Message: "Something went wrong",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		httpResponse.Success{
			Code:    http.StatusOK,
			Message: "User retrieved successfully",
			Data:    user,
		},
	)
}
