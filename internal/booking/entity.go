package booking

import (
	"time"

	"github.com/saurav11sarkar/ticket/internal/booking/dto"
	"github.com/saurav11sarkar/ticket/internal/event"
	"github.com/saurav11sarkar/ticket/internal/user"
)

type BookingEntity struct {
	ID      string      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID  string      `json:"user_id" gorm:"type:uuid;not null;index"`
	User    user.User   `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventID string      `json:"event_id" gorm:"type:uuid;not null;index"`
	Event   event.Event `json:"event" gorm:"foreignKey:EventID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Quantity   int `json:"quantity" gorm:"not null;check:quantity > 0"`
	TotalPrice int `json:"total_price" gorm:"not null;check:total_price > 0"`

	Status      string `json:"status" gorm:"type:varchar(50);not null;default:'pending'"`
	BookingCode string `json:"booking_code" gorm:"uniqueIndex;not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *BookingEntity) ToResponse() *dto.BookingResponse {
	return &dto.BookingResponse{
		ID:          b.ID,
		UserID:      b.UserID,
		EventID:     b.EventID,
		Quantity:    b.Quantity,
		TotalPrice:  b.TotalPrice,
		Status:      b.Status,
		BookingCode: b.BookingCode,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
