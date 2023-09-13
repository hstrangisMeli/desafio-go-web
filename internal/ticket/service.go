package tickets

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, dest string) (int, error)
	AverageDestination(ctx context.Context, dest string) (float64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTotalTickets(ctx context.Context, dest string) (int, error) {
	destTickets, err := s.r.GetTicketByDestination(ctx, dest)
	if err != nil {
		return 0, err
	}
	return len(destTickets), nil
}

func (s *service) AverageDestination(ctx context.Context, dest string) (float64, error) {
	var err error
	var totalTickets, destTickets []domain.Ticket
	totalTickets, err = s.r.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	destTickets, err = s.r.GetTicketByDestination(ctx, dest)
	if err != nil {
		return 0, err
	}
	nTotalTickets := len(totalTickets)
	nDestTickets := len(destTickets)
	return (float64(nDestTickets) / float64(nTotalTickets)) * 100, nil
}
