package main

import (
	"claseEnVivo1/ejercicios"
	"fmt"
)

func main() {
	res, err := ejercicios.SalaryTax(60000)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	maximum, err := ejercicios.Operation("maximum")
	minimum, err := ejercicios.Operation("minimum")
	average, err := ejercicios.Operation("average")
	_, err = ejercicios.Operation("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ejercicios.NoteAverage(1, 2, 3, 4, 5))
	fmt.Println(maximum(1, 2, 3, 4, 5))
	fmt.Println(minimum(1, 2, 3, 4, 5))
	fmt.Println(average(1, 2, 3, 4, 5))

	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)
	animalDog, err := ejercicios.Animal(perro)
	animalCat, err := ejercicios.Animal(gato)
	animalHamster, err := ejercicios.Animal(hamster)
	animalTarantula, err := ejercicios.Animal(tarantula)
	if err != nil {
		fmt.Println(err)
	}
	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(7)
	amount += animalTarantula(10)

	fmt.Print(amount)
}
