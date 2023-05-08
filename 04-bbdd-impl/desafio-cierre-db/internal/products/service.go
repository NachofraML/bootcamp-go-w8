package products

import (
	"encoding/json"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"os"
)

type Service interface {
	LoadFromJson() error
	Create(product *domain.Products) error
	ReadAll() ([]*domain.Products, error)
	GetTopProducts(limit int) ([]*domain.TopProducts, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) LoadFromJson() error {
	var products []*domain.Products

	file, err := os.Open("datos/products.json")
	if err != nil {
		return err
	}

	myDecoder := json.NewDecoder(file)
	if err = myDecoder.Decode(&products); err != nil {
		return err
	}

	for _, product := range products {
		_, err = s.r.Create(product)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Create(product *domain.Products) error {
	_, err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Products, error) {
	return s.r.ReadAll()
}

func (s *service) GetTopProducts(limit int) ([]*domain.TopProducts, error) {
	// Check limit before sending to repository
	if limit <= 0 || limit > 20 {
		//Here I could throw an error, but is more fast making this way now
		limit = 5
	}
	topProducts, err := s.r.GetTopProducts(limit)
	if err != nil {
		return nil, err
	}
	return topProducts, nil
}
