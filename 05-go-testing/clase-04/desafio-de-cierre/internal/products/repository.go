package products

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

var (
	ErrMockedRepo = errors.New("something went wrong")
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}

// RepositoryMock The repository of the upper code imitates a real code, so I need to create new mock to handle
// the responses better
type RepositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (r *RepositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := r.Called(sellerID)

	r0 := args.Get(0).([]Product)
	r1 := args.Error(1)

	return r0, r1
}
