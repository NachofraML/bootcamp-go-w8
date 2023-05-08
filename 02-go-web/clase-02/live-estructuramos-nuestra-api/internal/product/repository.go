package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/domain"
)

type Repository interface {
	GetByID(id int) (*domain.Producto, error)
	Save(producto *domain.Producto) (int, error)
}

var (
	ErrRepoInternal  = errors.New("internal error")
	ErrRepoNotUnique = errors.New("product already exists")
	ErrRepoNotFound  = errors.New("product not found")
)
