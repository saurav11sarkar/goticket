package dto

import "time"

type BookingResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	EventID     string    `json:"event_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"total_price"`
	Status      string    `json:"status"`
	BookingCode string    `json:"booking_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
