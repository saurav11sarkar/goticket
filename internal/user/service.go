package user

import (
	"context"
	"errors"
	"fmt"
	"html"
	"log"

	"github.com/saurav11sarkar/ticket/internal/auth"
	"github.com/saurav11sarkar/ticket/internal/common"
	"github.com/saurav11sarkar/ticket/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

const (
	userCreatedEmailSubject = "User created successfully"
	userCreatedEmailHTML    = `<h1>Welcome to Tricket, %s!</h1><p>Your account has been created successfully.</p>`
)

type EmailSender interface {
	SendEmail(email, subject, html string) error
}

type Service struct {
	jwtService  auth.JwtService
	repository  Repository
	emailSender EmailSender
}

func NewUserService(repository Repository, jwtService auth.JwtService, emailSender EmailSender) *Service {
	return &Service{
		repository:  repository,
		jwtService:  jwtService,
		emailSender: emailSender,
	}
}

func (s *Service) CreateUser(req dto.CreateUserRequest, profileImage string) (dto.CreateUserResponse, error) {
	user := User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     req.Password,
		ProfileImage: profileImage,
	}

	if err := s.repository.CreateUser(&user); err != nil {
		return dto.CreateUserResponse{}, err
	}

	emailHTML := fmt.Sprintf(userCreatedEmailHTML, html.EscapeString(user.Name))
	if err := s.emailSender.SendEmail(user.Email, userCreatedEmailSubject, emailHTML); err != nil {
		// The user is already persisted, so a welcome-email failure must not make
		// registration look unsuccessful to the client.
		log.Printf("send user creation email: %v", err)
	}

	return toUserResponse(&user), nil
}

func (s *Service) LoginUser(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("get user by email: %w", err)
	}
	if user == nil {
		return dto.LoginResponse{}, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return dto.LoginResponse{}, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Name:         user.Name,
		Email:        user.Email,
		Token:        token,
		ProfileImage: user.ProfileImage,
	}, nil
}

func (s *Service) GetMe(userID string) (dto.LoginResponse, error) {
	user, err := s.repository.GetUserById(userID)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("get user by ID: %w", err)
	}
	if user == nil {
		return dto.LoginResponse{}, ErrUserNotFound
	}

	return dto.LoginResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Name:         user.Name,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}, nil
}

func (s *Service) GetAll(
	ctx context.Context,
	query dto.UserQuery,
) ([]dto.UserResponse, common.Meta, error) {
	allowedSortFields := map[string]string{
		"name":      "name",
		"email":     "email",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
	}
	pagination := common.NewPagination(
		query.Page,
		query.Limit,
		query.SortBy,
		query.SortOrder,
		allowedSortFields,
	)

	users, total, err := s.repository.GetAll(ctx, query, pagination)
	if err != nil {
		return nil, common.Meta{}, fmt.Errorf("get users: %w", err)
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, toUserResponse(user))
	}

	return responses, common.NewMeta(total, pagination), nil
}

func toUserResponse(user *User) dto.UserResponse {
	return dto.UserResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Name:         user.Name,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}
}
