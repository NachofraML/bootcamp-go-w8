package products

import "meli-bootcamp/pkg/store"

type Product struct {

	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product


type Repository interface {
	Store(nombre string, tipo string, cantidad int, precio float64) (Product, error)

}
type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{nombre, tipo, cantidad, precio}
	r.db.Read(&ps)
	ps = append(ps, p)
	r.db.Write(ps)
	return p, nil
}