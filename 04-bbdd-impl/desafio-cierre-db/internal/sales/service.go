package sales

import (
	"encoding/json"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"os"
)

type Service interface {
	LoadFromJson() error
	Create(sales *domain.Sales) error
	ReadAll() ([]*domain.Sales, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) LoadFromJson() error {
	var sales []*domain.Sales

	file, err := os.Open("datos/sales.json")
	if err != nil {
		return err
	}

	myDecoder := json.NewDecoder(file)
	if err = myDecoder.Decode(&sales); err != nil {
		return err
	}

	for _, sale := range sales {
		_, err = s.r.Create(sale)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Create(sales *domain.Sales) error {
	_, err := s.r.Create(sales)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Sales, error) {
	return s.r.ReadAll()
}
