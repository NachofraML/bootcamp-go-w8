package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/domain"
)

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{&[]domain.Producto{}}
}

type InMemoryRepository struct {
	Productos *[]domain.Producto
}

func (repository *InMemoryRepository) GetByID(id int) (*domain.Producto, error) {
	for _, p := range *repository.Productos {
		if p.ID == id {
			return &p, nil
		}
	}
	return &domain.Producto{}, ErrRepoNotFound
}

func (repository *InMemoryRepository) Save(producto *domain.Producto) (int, error) {
	err := repository.validCodeValue(producto.CodeValue)
	if err != nil {
		return 0, ErrRepoNotUnique
	}
	producto.ID = len(*repository.Productos) + 1
	*repository.Productos = append(*repository.Productos, *producto)
	return producto.ID, nil
}

func (repository *InMemoryRepository) validCodeValue(actualCode string) error {
	for _, p := range *repository.Productos {
		if p.CodeValue == actualCode {
			return errors.New("cannot create code_value, unique violation")
		}
	}
	return nil
}
