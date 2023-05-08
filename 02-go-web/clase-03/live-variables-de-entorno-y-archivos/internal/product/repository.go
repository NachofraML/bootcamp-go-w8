package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/domain"
)

type Repository interface {
	GetByID(id int) (*domain.Producto, error)
	Save(producto *domain.Producto) (int, error)
	Update(id int, producto *domain.Producto) error
	Delete(id int) error
}

var (
	ErrRepoInternal  = errors.New("internal error")
	ErrRepoNotUnique = errors.New("product already exists")
	ErrRepoNotFound  = errors.New("product not found")
)
