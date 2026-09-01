package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/saurav11sarkar/ticket/internal/common"
	"github.com/saurav11sarkar/ticket/internal/user/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	GetAll(
		ctx context.Context,
		filter dto.UserQuery,
		pagination common.Pagination,
	) ([]*User, int64, error)
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

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetUserById(id string) (*User, error) {
	var user User
	result := r.db.Where(&User{ID: id}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) GetAll(
	ctx context.Context,
	filter dto.UserQuery,
	pagination common.Pagination,
) ([]*User, int64, error) {
	users := make([]*User, 0)
	var total int64

	query := r.db.
		WithContext(ctx).
		Model(&User{}).
		Scopes(common.SearchScope(filter.SearchTerm, "name", "email"))

	query = applyUserFilters(query, filter)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count users: %w", err)
	}

	if err := query.
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: pagination.SortBy},
			Desc:   pagination.SortOrder == "desc",
		}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
		}).
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Find(&users).
		Error; err != nil {
		return nil, 0, fmt.Errorf("find users: %w", err)
	}

	return users, total, nil
}

func applyUserFilters(db *gorm.DB, filter dto.UserQuery) *gorm.DB {
	if name := strings.TrimSpace(filter.Name); name != "" {
		db = db.Where("name = ?", name)
	}
	if email := strings.TrimSpace(filter.Email); email != "" {
		db = db.Where("email = ?", email)
	}

	return db
}
