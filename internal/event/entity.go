package event

import (
	"time"

	"github.com/saurav11sarkar/ticket/internal/event/dto"
)

type Event struct {
	ID string `json:"id" gorm:"uuid,default:gen_random_uuid();primary_key"`
	//UserID          string    `json:"user_id" gorm:"type:uuid,not null"`
	//User            user.User `json:"user" gorm:"foreignkey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title           string    `json:"title" gorm:"type:text not null"`
	Description     string    `json:"description" gorm:"type:text"`
	Location        string    `json:"location" gorm:"type:text not null"`
	StartAt         time.Time `json:"start_at" gorm:"type:timestamp not null"`
	TotalTicket     int       `json:"total_ticket" gorm:"default:0"`
	AvailableTicket int       `json:"available_ticket" gorm:"default:0"`
	Price           int       `json:"price" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (e *Event) ToResponse() *dto.EventResponse {
	return &dto.EventResponse{
		ID:              e.ID,
		Title:           e.Title,
		Description:     e.Description,
		Location:        e.Location,
		StartAt:         e.StartAt,
		TotalTicket:     e.TotalTicket,
		AvailableTicket: e.AvailableTicket,
		Price:           e.Price,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}
