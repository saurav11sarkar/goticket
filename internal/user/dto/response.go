package dto

import "time"

type UserResponse struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	ProfileImage string    `json:"profile_image"`
}

type CreateUserResponse = UserResponse

type LoginResponse struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Token        string    `json:"token,omitempty"`
	ProfileImage string    `json:"profile_image"`
}
