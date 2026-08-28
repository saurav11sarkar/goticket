package event

import "github.com/saurav11sarkar/ticket/internal/event/dto"

type Service interface {
	Create(req dto.CreateEventRequest) (*dto.EventResponse, error)
	GetAll() ([]*dto.EventResponse, error)
	GetByID(eventID string) (*dto.EventResponse, error)
	Update(eventID string, req dto.UpdateEventRequest) (*dto.EventResponse, error)
	Delete(eventID string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(req dto.CreateEventRequest) (*dto.EventResponse, error) {
	event := Event{
		Title:           req.Title,
		Description:     req.Description,
		Location:        req.Location,
		StartAt:         req.StartAt,
		TotalTicket:     req.TotalTicket,
		AvailableTicket: req.TotalTicket,
		Price:           req.Price,
	}
	if err := s.repo.Create(&event); err != nil {
		return nil, err
	}

	//return &dto.EventResponse{
	//	ID:              event.ID,
	//	Title:           event.Title,
	//	Description:     event.Description,
	//	StartAt:         event.StartAt,
	//	TotalTicket:     event.TotalTicket,
	//	AvailableTicket: event.AvailableTicket,
	//	Price:           event.Price,
	//	CreatedAt:       event.CreatedAt,
	//	UpdatedAt:       event.UpdatedAt,
	//}, nil

	return event.ToResponse(), nil
}

func (s *service) GetAll() ([]*dto.EventResponse, error) {
	// implementation
	return nil, nil
}

func (s *service) GetByID(eventID string) (*dto.EventResponse, error) {
	// implementation
	return nil, nil
}

func (s *service) Update(
	eventID string,
	req dto.UpdateEventRequest,
) (*dto.EventResponse, error) {
	// implementation
	return nil, nil
}

func (s *service) Delete(eventID string) error {
	// implementation
	return nil
}
