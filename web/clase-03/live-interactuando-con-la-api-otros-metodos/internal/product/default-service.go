package product

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-interactuando-con-la-api-otros-metodos/internal/domain"
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

func (service *DefaultService) Update(id int, producto *domain.Producto) (err error) {
	err = producto.Validate()
	if err != nil {
		err = ErrServiceInvalid
		return
	}
	err = service.Repository.Update(id, producto)
	if err != nil {
		if errors.Is(err, ErrRepoNotUnique) {
			err = ErrServiceNotUnique
			return
		}
		if errors.Is(err, ErrRepoNotFound) {
			err = ErrServiceNotFound
			return
		}
		err = ErrServiceInternal
		return
	}
	return
}

func (service *DefaultService) Delete(id int) (err error) {
	err = service.Repository.Delete(id)
	if err != nil {
		if errors.Is(err, ErrRepoNotFound) {
			err = ErrServiceNotFound
			return
		}
		err = ErrServiceInternal
		return
	}
	return
}
