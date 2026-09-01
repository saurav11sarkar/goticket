package dto

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,min=10"`
	Location    string    `json:"location" validate:"required"`
	StartAt     time.Time `json:"start_at" validate:"required"`
	TotalTicket int       `json:"total_ticket" validate:"required,gt=0"`
	Price       int       `json:"price" validate:"required,gt=0"`
}

type UpdateEventRequest struct {
	Title       *string    `json:"title" validate:"omitempty,max=255"`
	Description *string    `json:"description" validate:"omitempty,min=10"`
	Location    *string    `json:"location" validate:"omitempty"`
	StartAt     *time.Time `json:"start_at" validate:"omitempty"`
	TotalTicket *int       `json:"total_ticket" validate:"omitempty,gt=0"`
	Price       *int       `json:"price" validate:"omitempty,gt=0"`
}

func (r UpdateEventRequest) IsEmpty() bool {
	return r.Title == nil &&
		r.Description == nil &&
		r.Location == nil &&
		r.StartAt == nil &&
		r.TotalTicket == nil &&
		r.Price == nil
}
