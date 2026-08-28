package dto

import "time"

type EventResponse struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	StartAt         time.Time `json:"start_at"`
	TotalTicket     int       `json:"total_ticket" `
	AvailableTicket int       `json:"available_ticket" `
	Price           int       `json:"price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
