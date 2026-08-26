package user

import (
	"errors"
	"fmt"

	"github.com/saurav11sarkar/ticket/internal/auth"
	"github.com/saurav11sarkar/ticket/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type Service struct {
	jwtService auth.JwtService
	repository Repository
}

func NewUserService(repository Repository, jwtService auth.JwtService) *Service {
	return &Service{repository: repository, jwtService: jwtService}
}

func (s *Service) CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.repository.CreateUser(&user); err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}, nil
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
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
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
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}, nil
}
