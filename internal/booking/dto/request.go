package dto

type BookingRequest struct {
	EventID  string `json:"event_id" validate:"required,uuid"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
}
