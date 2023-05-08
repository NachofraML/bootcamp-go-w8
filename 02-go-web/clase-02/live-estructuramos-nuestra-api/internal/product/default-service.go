package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/domain"
)

func NewDefaultService(rp Repository) Service {
	return &DefaultService{Repository: rp}
}

type DefaultService struct {
	Repository Repository
}

func (service *DefaultService) GetByID(id int) (productoEncontrado *domain.Producto, err error) {
	productoEncontrado, err = service.Repository.GetByID(id)
	if err != nil {
		return
	}
	if productoEncontrado == nil {
		err = ErrServiceNotFound
		return
	}
	return
}
func (service *DefaultService) Save(producto *domain.Producto) (err error) {
	err = producto.Validate()
	if err != nil {
		err = ErrServiceInvalid
		return
	}
	_, err = service.Repository.Save(producto)
	if err != nil {
		if errors.Is(err, ErrRepoNotUnique) {
			err = ErrServiceNotUnique
			return
		}
		err = ErrServiceInternal
		return
	}
	return
}
