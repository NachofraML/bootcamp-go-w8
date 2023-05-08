package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/domain"
)

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{[]*domain.Producto{}}
}

type InMemoryRepository struct {
	Productos []*domain.Producto
}

func (repository *InMemoryRepository) GetByID(id int) (*domain.Producto, error) {
	for _, p := range repository.Productos {
		if p.ID == id {
			return p, nil
		}
	}
	return &domain.Producto{}, ErrRepoNotFound
}

func (repository *InMemoryRepository) Save(producto *domain.Producto) (int, error) {
	err := repository.validCodeValue(producto.CodeValue)
	if err != nil {
		return 0, ErrRepoNotUnique
	}
	producto.ID = len(repository.Productos) + 1
	repository.Productos = append(repository.Productos, producto)
	return producto.ID, nil
}

func (repository *InMemoryRepository) Update(id int, producto *domain.Producto) (err error) {
	producto.ID = id

	for i, p := range repository.Productos {
		if p.ID == producto.ID {
			// Si pasaron el CodeValue mismo que existe actualmente para hacer PATCH,
			// no lo verifico porque es imposible que genere conflictos
			if producto.CodeValue == repository.Productos[i].CodeValue {
				repository.Productos[i] = producto
				return
			} else {
				// En caso de que hayan pasado un CodeValue diferente, ahora si lo verifico para que no se repita
				// en ningun producto
				if err = repository.validCodeValue(producto.CodeValue); err != nil {
					return ErrRepoNotUnique
				}
			}
			repository.Productos[i] = producto
			return
		}
	}
	err = ErrRepoNotFound
	return
}

func (repository *InMemoryRepository) Delete(id int) (err error) {
	for i, p := range repository.Productos {
		if p.ID == id {
			repository.Productos = append(repository.Productos[:i], repository.Productos[i+1:]...)
			return
		}
	}
	err = ErrRepoNotFound
	return
}

func (repository *InMemoryRepository) validCodeValue(actualCode string) error {
	for _, p := range repository.Productos {
		if p.CodeValue == actualCode {
			return errors.New("cannot create code_value, unique violation")
		}
	}
	return nil
}
