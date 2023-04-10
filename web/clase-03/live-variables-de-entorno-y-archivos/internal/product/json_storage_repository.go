package product

import (
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/domain"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/pkg/store"
)

func NewJsonStorageRepository() Repository {
	jsonStorage, err := store.NewJsonStorage("./web/clase-03/live-variables-de-entorno-y-archivos/pkg/store/productos.json")
	if err != nil {
		panic(err)
	}
	return &JsonStorageRepository{jsonStorage}
}

type JsonStorageRepository struct {
	jsonStorage *store.JsonStorage
}

func (sr JsonStorageRepository) GetByID(id int) (producto *domain.Producto, err error) {
	producto, err = sr.jsonStorage.GetProductById(id)
	if err != nil {
		return
	}
	return
}

func (sr JsonStorageRepository) Save(producto *domain.Producto) (lastInsertId int, err error) {
	lastInsertId, err = sr.jsonStorage.CreateProduct(producto)
	if err != nil {
		return
	}
	return
}

func (sr JsonStorageRepository) Update(id int, producto *domain.Producto) (err error) {
	err = sr.jsonStorage.UpdateProduct(id, producto)
	if err != nil {
		fmt.Printf("repository json: %e", err)
		return
	}
	return
}

func (sr JsonStorageRepository) Delete(id int) (err error) {
	err = sr.jsonStorage.DeleteProduct(id)
	if err != nil {
		return
	}
	return
}
