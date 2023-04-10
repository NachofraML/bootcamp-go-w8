package product

import (
	"errors"
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/domain"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/pkg/store"
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
		if errors.Is(err, ErrRepoNotFound) || errors.Is(err, store.ErrJsonStorageProductNotFound) {
			err = ErrServiceNotFound
			return
		}
		fmt.Println(err)
		err = ErrRepoInternal
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
		fmt.Println(err)
		err = ErrServiceInternal
		return
	}
	return
}

func (service *DefaultService) Update(id int, producto *domain.Producto) (err error) {
	if producto.Expiration != "" {
		err = producto.Validate()
		if err != nil {
			err = ErrServiceInvalid
			return
		}
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
		fmt.Printf("SERVICE: %e", err)
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
		fmt.Println(err)
		err = ErrServiceInternal
		return
	}
	return
}
