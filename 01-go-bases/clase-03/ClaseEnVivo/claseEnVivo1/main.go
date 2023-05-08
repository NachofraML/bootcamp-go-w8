package main

import (
	"errors"
)

// Product 1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Products 2. Tener un slice global de Product llamado Products instanciado con valores.

type Products struct {
	Products []Product
}

// 3. 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save()
// deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.

func (product Product) Save(products *Products) {
	products.Products = append(products.Products, product)
}

// El método GetAll() deberá imprimir todos los productos guardados en el slice Products.

func (product Product) getAll(products Products) {
	for _, p := range products.Products {
		println(p.Name)
	}
}

//  4. Una función getById() al cual se le deberá pasar un
//     INT como parámetro y retorna el producto correspondiente al parámetro pasado.

func (products Products) getById(id int) (Product, error) {
	for _, p := range products.Products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, errors.New("The product does not exist")
}

// 1.5. Ejecutar al menos una vez cada método y función definido desde ex().

// 2.4. Instanciar en la función ex() tanto una Person como un Employee cargando sus
// respectivos campos y por último ejecutar el método PrintEmployee().

func main() {
	products := Products{}

	product := Product{
		ID:          1,
		Name:        "Juan",
		Price:       10.00,
		Description: "Juan",
		Category:    "Juan",
	}
	product2 := Product{
		ID:          2,
		Name:        "Juan",
		Price:       10.00,
		Description: "Juan",
		Category:    "Juan",
	}

	product.Save(&products)
	product2.Save(&products)

	println(products.Products)

	product.getAll(products)

	obtainedProduct, err := products.getById(1)
	println("Finded product name:", obtainedProduct.Name, "Error: ", err)

	person := Person{1, "Nacho", "10-10-1980"}
	employee := Employee{1, "Software Developer", person}

	employee.PrintEmployee()
}

// Ejercicio 2 - Employee
// Crear una estructura Person con los campos ID, Name, DateOfBirth.

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

// Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.

type Employee struct {
	ID       int
	Position string
	Person
}

// Realizar el método a la estructura Employe que se llame PrintEmployee(),
// lo que hará es realizar la impresión de los campos de un empleado.

func (e *Employee) PrintEmployee() {
	println(e.ID, e.Name, e.DateOfBirth)
}
