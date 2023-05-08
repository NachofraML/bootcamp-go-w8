package invoices

import (
	"encoding/json"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"os"
	"time"
)

type Service interface {
	LoadFromJson() error
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	Update(invoices *domain.Invoices) error
	UpdateAllTotals() error
	CalculateTotal(id int) (float64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) LoadFromJson() error {
	var invoices []*domain.Invoices

	file, err := os.Open("datos/invoices.json")
	if err != nil {
		return err
	}

	myDecoder := json.NewDecoder(file)
	if err = myDecoder.Decode(&invoices); err != nil {
		return err
	}

	for _, invoice := range invoices {
		_, err = s.r.Create(invoice)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) ReadAll() ([]*domain.Invoices, error) {
	return s.r.ReadAll()
}

func (s *service) Update(invoices *domain.Invoices) error {
	err := s.r.Update(invoices)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateAllTotals() error {
	var total float64
	invoices, err := s.ReadAll()
	if err != nil {
		return err
	}
	for _, invoice := range invoices {
		total, err = s.CalculateTotal(invoice.Id)
		if err != nil {
			return err
		}
		invoice.Total = total

		datetime, err := time.Parse(time.RFC3339, invoice.Datetime)
		if err != nil {
			return err
		}
		datetimeFormatted := datetime.Format("2006-01-02 15:04:05")
		invoice.Datetime = datetimeFormatted

		err = s.Update(invoice)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) CalculateTotal(id int) (float64, error) {
	var total float64
	invoiceProducts, err := s.r.GetInvoiceProducts(id)
	if err != nil {
		return 0, err
	}
	for _, invoiceProduct := range invoiceProducts {
		total += invoiceProduct.Price * float64(invoiceProduct.Quantity)
	}
	return total, nil
}
