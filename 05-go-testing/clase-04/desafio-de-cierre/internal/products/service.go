package products

import (
	"github.com/stretchr/testify/mock"
	"log"
)

type Service interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllBySeller(sellerID string) ([]Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)
	if err != nil {
		log.Println("error in repository", err.Error(), "sellerId:", sellerID)
		return nil, err
	}
	return data, err
}

// ServiceMock for Handler unit testing
type ServiceMock struct {
	mock.Mock
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (s *ServiceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := s.Called(sellerID)

	r0 := args.Get(0).([]Product)
	r1 := args.Error(1)

	return r0, r1
}
