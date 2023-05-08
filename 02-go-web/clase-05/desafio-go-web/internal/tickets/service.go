package tickets

import (
	"context"
	"errors"
	"strings"
)

type Service interface {
	GetTotalTickets(ctx context.Context, country string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	db Repository
}

func NewService(repo Repository) Service {
	return &service{
		db: repo,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	destination = strings.ToUpper(string(destination[0])) + strings.ToLower(destination[1:])
	tickets, err := s.db.GetTicketByDestination(ctx, destination)
	if err != nil {
		if errors.Is(err, ErrRepositoryTicketsDbEmpty) {
			err = ErrServiceTicketsDbEmpty
		}
		return 0, err
	}
	if len(tickets) == 0 {
		err = ErrServiceTicketsNotFound
	}
	return len(tickets), err
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	destination = strings.ToUpper(string(destination[0])) + strings.ToLower(destination[1:])
	var totalTicketsPerCountry float64
	tickets, err := s.db.GetAll(ctx)
	if err != nil {
		if errors.Is(err, ErrRepositoryTicketsDbEmpty) {
			err = ErrServiceTicketsDbEmpty
		}
		return 0, err
	}

	for _, ticket := range tickets {
		if ticket.Country == destination {
			totalTicketsPerCountry++
		}
	}
	if totalTicketsPerCountry == 0 {
		err = ErrServiceTicketsNotFound
		return 0, err
	}

	result := totalTicketsPerCountry / float64(len(tickets)) * 100.0
	return result, nil
}

var (
	ErrServiceTicketsDbEmpty  = errors.New("database is empty")
	ErrServiceTicketsNotFound = errors.New("empty list of tickets")
)
