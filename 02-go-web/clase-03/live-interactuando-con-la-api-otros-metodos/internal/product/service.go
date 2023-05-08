package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-interactuando-con-la-api-otros-metodos/internal/domain"
)

type Service interface {
	GetByID(id int) (*domain.Producto, error)
	Save(producto *domain.Producto) error
	Update(id int, producto *domain.Producto) error
	Delete(id int) error
}

var (
	ErrServiceInternal  = errors.New("internal error")
	ErrServiceInvalid   = errors.New("invalid product")
	ErrServiceNotUnique = errors.New("product already exists")
	ErrServiceNotFound  = errors.New("product not found")
)
