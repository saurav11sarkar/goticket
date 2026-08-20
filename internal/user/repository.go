package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type Repository interface {
	CreateUser(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error == nil {
		return nil
	}
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return ErrUserAlreadyExists
	}

	return result.Error
}
