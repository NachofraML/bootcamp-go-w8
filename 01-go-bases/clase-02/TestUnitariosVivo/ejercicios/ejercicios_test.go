package ejercicios

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEjercicios(t *testing.T) {

	//Ejercicio 1 - Testear el impuesto del salario
	//La empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, al momento de
	//depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos.
	//Para esto nos encargaron el trabajo de realizar los test correspondientes para:
	//Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
	//Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
	//Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
	t.Run("SalaryTax", func(t *testing.T) {
		//Given
		var err error
		//When

		resultLessThan50000, err := SalaryTax(25000)
		resultMoreThan50000, err := SalaryTax(75000)
		resultMoreThan150000, err := SalaryTax(200000)
		_, err = SalaryTax(0)

		expectedLessThan50000 := 4250.0
		expectedMoreThan50000 := 12750.0
		expectedMoreThan150000 := 54000.0

		//Then
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf("%.6f", resultLessThan50000), fmt.Sprintf("%.6f", expectedLessThan50000))
		assert.Equal(t, fmt.Sprintf("%.6f", resultMoreThan50000), fmt.Sprintf("%.6f", expectedMoreThan50000))
		assert.Equal(t, fmt.Sprintf("%.6f", resultMoreThan150000), fmt.Sprintf("%.6f", expectedMoreThan150000))
	})

	//Ejercicio 2 - Calcular promedio
	//El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente,
	//por lo que ahora nos corresponde realizar los test correspondientes:
	//Calcular el promedio de las notas de los alumnos.
	t.Run("TestNoteAverage", func(t *testing.T) {
		//Given
		var notes = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 5, -10}
		expectedAverage := 5.0

		//When
		resultAverage := NoteAverage(notes...)

		//Then
		assert.Equal(t, resultAverage, expectedAverage, "Salary of category A must be 1000.0")
	})

	//Ejercicio 3 - Test del salario
	//La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios,
	//por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes
	//pruebas en nuestro código:
	//Calcular el salario de la categoría “A”.
	//Calcular el salario de la categoría “B”.
	//Calcular el salario de la categoría “C”.
	t.Run("TestCalcSalary", func(t *testing.T) {
		//Given
		var minutes = 60.0
		var err error
		expectedCategoryA := 1000.0
		expectedCategoryB := 1800.0
		expectedCategoryC := 4500.0

		//When
		resultCategoryA, err := CalcSalary(minutes, "A")
		resultCategoryB, err := CalcSalary(minutes, "B")
		resultCategoryC, err := CalcSalary(minutes, "C")
		_, err = CalcSalary(minutes, "D")

		//Then
		assert.Error(t, err)
		assert.Equal(t, resultCategoryA, expectedCategoryA, "Salary of category A must be 1000.0")
		assert.Equal(t, resultCategoryB, expectedCategoryB, "Salary of category B must be 1800.0")
		assert.Equal(t, resultCategoryC, expectedCategoryC, "Salary of category C must be 4500.0")
	})

	//Ejercicio 4 - Testear el cálculo de estadísticas
	//Los profesores de la universidad de Colombia, entraron al programa de análisis de datos de Google,
	//el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar
	//el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
	//Realizar test para calcular el mínimo de calificaciones.
	//Realizar test para calcular el máximo de calificaciones.
	//Realizar test para calcular el promedio de calificaciones.

	t.Run("TestOperation", func(t *testing.T) {
		//Given
		var err error
		toTestInts := []int{1, 2, 3, 4, 5}
		expectedMinimum := 1.0
		expectedMaximum := 5.0
		expectedAverage := 3.0

		//When
		minimum, err := Operation("minimum")
		maximum, err := Operation("maximum")
		average, err := Operation("average")
		resultMinimum := minimum(toTestInts...)
		resultMaximum := maximum(toTestInts...)
		resultAverage := average(toTestInts...)
		_, err = Operation("nothing")

		//Then
		assert.Error(t, err)
		assert.Equal(t, resultMinimum, expectedMinimum)
		assert.Equal(t, resultMaximum, expectedMaximum)
		assert.Equal(t, resultAverage, expectedAverage)
	})

	//Ejercicio 5 - Calcular cantidad de alimento
	//El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y
	//compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar
	//los test necesarios para verificar que todo funcione correctamente:
	//Verificar el cálculo de la cantidad de alimento para los perros.
	//Verificar el cálculo de la cantidad de alimento para los gatos.
	//Verificar el cálculo de la cantidad de alimento para los hamster.
	//Verificar el cálculo de la cantidad de alimento para las tarántulas.

	t.Run("TestAnimal", func(t *testing.T) {
		//Given
		var err error
		var animalCount = 10.0
		expectedPerro := 100.0
		expectedGato := 50.0
		expectedHamster := 2.5
		expectedTarantula := 1.5

		//When
		perro, err := Animal("perro")
		gato, err := Animal("gato")
		hamster, err := Animal("hamster")
		tarantula, err := Animal("tarantula")
		resultPerro := perro(animalCount)
		resultGato := gato(animalCount)
		resultHamster := hamster(animalCount)
		resultTarantula := tarantula(animalCount)
		_, err = Animal("nothing")

		//Then
		assert.Error(t, err)
		assert.Equal(t, resultPerro, expectedPerro)
		assert.Equal(t, resultGato, expectedGato)
		assert.Equal(t, resultHamster, expectedHamster)
		assert.Equal(t, resultTarantula, expectedTarantula)
	})
}
