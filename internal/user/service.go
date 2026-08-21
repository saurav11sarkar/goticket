package user

import (
	"github.com/saurav11sarkar/ticket/internal/user/dto"
)

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
