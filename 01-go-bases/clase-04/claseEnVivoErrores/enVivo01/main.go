package main

import (
	"errors"
	"fmt"
)

type ErrorSalario struct {
	Mensaje string
}

func (e *ErrorSalario) Error() string {
	return e.Mensaje
}

var (
	ErrTaxSalary      = &ErrorSalario{Mensaje: "Error: el mínimo imponible es de 150.000 y el salario ingresado es de: "}
	ErrSalaryQuantity = &ErrorSalario{Mensaje: "Error: el salario es menor a 10.000"}
	ErrWorkedHours    = &ErrorSalario{Mensaje: "Error: el trabajador no puede haber trabajado menos de 80 hs mensuales"}
	ErrValuePerHour   = &ErrorSalario{Mensaje: "Error: el salario es menor o igual a 0"}
)

//Ejercicio 1 - Impuestos de salario #1
//En tu función “ex”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
//Creá un error personalizado con un struct que implemente “Error()” con el mensaje
//“Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000.
//De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.

func aplicaImpuesto(salario int) (string, error) {
	if salario < 150000 {
		return "", ErrTaxSalary
	}
	return "Debe pagar impuesto", nil
}

//Ejercicio 2 - Impuestos de salario #2
//En tu función “ex”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
//Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor a 10.000"
//y lanzalo en caso de que “salary” sea menor o igual a 10.000.
//La validación debe ser hecha con la función “Is()” dentro del “ex”.

func validarMontoSalario(salario int) error {
	if salario <= 10000 {
		return ErrSalaryQuantity
	}
	return nil
}

func main() {
	salario := 100
	_, errImpuesto := aplicaImpuesto(salario)
	errMontoSalario := validarMontoSalario(salario)
	salarioTotal, _ := calcularSalarioMensualPorHorasTrabajadas(80, 1000)
	_, errWorkedHours := calcularSalarioMensualPorHorasTrabajadas(0, 1000)
	_, errValuePerHour := calcularSalarioMensualPorHorasTrabajadas(80, 0)

	if errors.Is(errImpuesto, ErrTaxSalary) {
		fmt.Printf(errImpuesto.Error()+"%v\n", salario)
	}
	if errors.Is(errMontoSalario, ErrSalaryQuantity) {
		fmt.Println(errMontoSalario.Error())
	}
	if errors.Is(errWorkedHours, ErrWorkedHours) {
		fmt.Println(errWorkedHours.Error())
	}
	if errors.Is(errValuePerHour, ErrValuePerHour) {
		fmt.Println(errValuePerHour.Error())
	}

	fmt.Println(salarioTotal)
}

//Ejercicio 3 - Impuestos de salario #3
//Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,
//en reemplazo de “Error()”,  se implemente “errors.New()”.

//Ejercicio 4 - Impuestos de salario #4
//Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por
//parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá
//decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor
//de tipo int pasado por parámetro).

// ESTOS DOS NO LOS HAGO PORQUE ME PARECEN ROMPEN TODO EL ESQUEMA DE VARIABLES, PERO LO QUE DEBERIA HACER ES TRABAJAR
// CON COPIAS EN VEZ DE PUNTEROS Y EN LAS FUNCIONES RETORNAR fmt.Error() JUNTO A LOS SALARIOS, Y DESPUES COMPARARLOS
// CON COPIAS QUE CREE EN MAIN PARA PODER TIRARLOS

//Ejercicio 5 -  Impuestos de salario #5
//Vamos a hacer que nuestro programa sea un poco más complejo y útil.
//Desarrollá las funciones necesarias para permitir a la empresa calcular:
//Salario mensual de un trabajador según la cantidad de horas trabajadas.
//La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
//Dicha función deberá retornar más de un valor (salario calculado y error).
//En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto
//de impuesto.
//En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función
//debe devolver un error. El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80
//hs mensuales”.

func calcularSalarioMensualPorHorasTrabajadas(horasTrabajadas float64, valorPorHora float64) (salarioCalculado float64, err error) {
	if valorPorHora <= 0 {
		err = ErrValuePerHour
		return
	}
	if horasTrabajadas < 80 {
		err = ErrWorkedHours
		return
	}
	salarioCalculado = valorPorHora * horasTrabajadas
	if salarioCalculado >= 150000 {
		salarioCalculado *= 0.90
	}
	return

}
