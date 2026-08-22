package user

import (
	"errors"
	"fmt"

	"github.com/saurav11sarkar/ticket/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type Service struct {
	repository Repository
}

func NewUserService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return dto.CreateUserResponse{}, err
	//}

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

	return dto.LoginResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}, nil
}
