package event

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrDuplicatedTicket = errors.New("duplicate ticket")
	ErrNotFound         = errors.New("not found")
)

type Repository interface {
	Create(event *Event) error
	GetAll() ([]*Event, error)
	GetByID(eventID string) (*Event, error)
	Update(eventID string, event *Event) error
	Delete(eventID string) error
}

type respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &respository{
		db: db,
	}
}

func (r *respository) Create(event *Event) error {
	result := r.db.Create(event)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrDuplicatedTicket
		}
		return result.Error
	}
	return nil
}
func (r *respository) GetAll() ([]*Event, error) {
	var events []*Event
	result := r.db.Find(&events)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []*Event{}, ErrNotFound
		}
		return []*Event{}, result.Error
	}
	return events, nil
}

func (r *respository) GetByID(eventID string) (*Event, error) {
	var event *Event
	result := r.db.Where(&Event{ID: eventID}).First(&event)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return event, nil
}

func (r *respository) Update(eventID string, event *Event) error {
	result := r.db.Model(&Event{}).Where(&Event{ID: eventID}).Updates(event)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *respository) Delete(eventID string) error {
	result := r.db.Where(&Event{ID: eventID}).Delete(&Event{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
