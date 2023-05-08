package customers

import (
	"encoding/json"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"os"
)

type Service interface {
	LoadFromJson() error
	Create(customers *domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
	GetConditionsTotals() ([]*domain.CustomerConditionsTotals, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) LoadFromJson() error {
	var customers []*domain.Customers

	file, err := os.Open("datos/customers.json")
	if err != nil {
		return err
	}

	myDecoder := json.NewDecoder(file)
	if err = myDecoder.Decode(&customers); err != nil {
		return err
	}

	for _, customer := range customers {
		_, err = s.r.Create(customer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Create(customers *domain.Customers) error {
	_, err := s.r.Create(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}

func (s *service) GetConditionsTotals() ([]*domain.CustomerConditionsTotals, error) {
	customerConditionsTotals, err := s.r.GetConditionsTotals()
	if err != nil {
		return nil, err
	}
	return customerConditionsTotals, nil
}
