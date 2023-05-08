package main

import "fmt"

type Alumnos struct {
	DNI            int
	Nombre         string
	Apellido       string
	FechaDeIngreso string
}

const (
	pequenio = "pequenio"
	mediano  = "mediano"
	grande   = "grande"
)

type Producto interface {
	calcularPrecio() float64
}

type ProductoPequenio struct {
	Precio float64
}

func (p ProductoPequenio) calcularPrecio() float64 {
	return p.Precio
}

type ProductoMediano struct {
	Precio float64
}

func (p ProductoMediano) calcularPrecio() float64 {
	return p.Precio + p.Precio*0.03
}

type ProductoGrande struct {
	Precio float64
}

func (p ProductoGrande) calcularPrecio() float64 {
	return p.Precio + p.Precio*0.06 + 2500
}

func productoFactory(tipoProducto string, precio float64) Producto {

	switch tipoProducto {
	case pequenio:
		return ProductoPequenio{Precio: precio}
	case mediano:
		return ProductoMediano{Precio: precio}
	case grande:
		return ProductoGrande{Precio: precio}
	default:
		return nil
	}

}

func (alumno Alumnos) detalle() {
	fmt.Println(alumno.DNI)
	fmt.Println(alumno.Nombre)
	fmt.Println(alumno.Apellido)
	fmt.Println(alumno.FechaDeIngreso)
}

func main() {
	alumno := Alumnos{DNI: 34465674, Nombre: "Nacho", Apellido: "Mosca", FechaDeIngreso: "10-10-2015"}
	alumno.detalle()

	productoPequenio := productoFactory(pequenio, 10000)
	productoMediano := productoFactory(mediano, 10000)
	productoGrande := productoFactory(grande, 10000)
	fmt.Println(productoPequenio.calcularPrecio())
	fmt.Println(productoMediano.calcularPrecio())
	fmt.Println(productoGrande.calcularPrecio())
}
