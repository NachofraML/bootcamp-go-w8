package ejercicios

import (
	"errors"
	"fmt"
	"math"
)

//Ejercicio 1 - Impuestos de salario
//Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para
//cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
//Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de
//$150.000 se le descontará además un 10 % (27% en total).

func SalaryTax(salary float64) (float64, error) {
	if salary <= 0 {
		return 0, errors.New("Salary must be positive and greater than 0")
	} else if salary > 0 && salary < 150000 {
		return salary * 0.17, nil
	}
	return salary * 0.27, nil
}

//Ejercicio 2 - Calcular promedio
//Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función
//en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.

func NoteAverage(values ...int) float64 {
	var res int
	for _, v := range values {
		if v < 0 {
			fmt.Printf("Ignoring negative number %v\n", v)
			continue
		}
		res = res + v
	}
	return float64(res) / float64(len(values))
}

//Ejercicio 3 - Calcular salario
//Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas
//por mes y la categoría.
//Categoría C, su salario es de $1.000 por hora.
//Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
//Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría
//y que devuelva su salario.

func CalcSalary(minutes float64, category string) (float64, error) {
	var salaryPerHour, finalSalary float64
	switch category {
	case "A":
		salaryPerHour = 1000.00
		finalSalary = salaryPerHour * (minutes / 60)
	case "B":
		salaryPerHour = 1500.00
		finalSalary += (salaryPerHour * (minutes / 60)) * 1.20
	case "C":
		salaryPerHour = 3000.00
		finalSalary += (salaryPerHour * (minutes / 60)) * 1.50
	default:
		return 0, errors.New("Category must be A, B or C ")
	}
	return finalSalary, nil
}

//Ejercicio 4 - Calcular estadísticas
//
//Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
//de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
//Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar
//(mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido)
//que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.

func Operation(operation string) (func(values ...int) float64, error) {
	switch operation {
	case "minimum":
		return Min, nil
	case "average":
		return NoteAverage, nil
	case "maximum":
		return Max, nil
	default:
		return nil, errors.New("unknown operation")
	}
}

func Min(values ...int) float64 {
	min := math.MaxFloat64
	for _, v := range values {
		if float64(v) < min {
			min = float64(v)
		}
	}
	return min
}

func Max(values ...int) (max float64) {
	for _, v := range values {
		if float64(v) > max {
			max = float64(v)
		}
	}
	return max
}

//Ejercicio 5 - Calcular cantidad de alimento
//
//Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
//Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a
//muchos animales más.
//Tienen los siguientes datos:
//Perro 10 kg de alimento.
//Gato 5 kg de alimento.
//Hamster 250 g de alimento.
//Tarántula 150 g de alimento.
//Se solicita:
//Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y
//que retorne una función y un mensaje (en caso que no exista el animal).
//Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal
//especificado.

func Animal(animal string) (func(float64) float64, error) {
	switch animal {
	case "perro":
		return Perro, nil
	case "gato":
		return Gato, nil
	case "hamster":
		return Hamster, nil
	case "tarantula":
		return Tarantula, nil
	default:
		return nil, errors.New("unknown animal")
	}
}

func Perro(animalAmount float64) float64 {
	return animalAmount * 10
}

func Gato(animalAmount float64) float64 {
	return animalAmount * 5
}

func Hamster(animalAmount float64) float64 {
	return animalAmount * 0.25
}

func Tarantula(animalAmount float64) float64 {
	return animalAmount * 0.15
}
