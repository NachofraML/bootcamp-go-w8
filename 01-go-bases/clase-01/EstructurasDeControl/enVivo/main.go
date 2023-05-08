package main

import "fmt"

func main() {
	//Ejercicio 1 - Letras de una palabra
	//La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por
	//separado para deletrearla. Para eso tendrán que:
	//Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
	//Luego, imprimí cada una de las letras.
	palabra := "Buenas"
	fmt.Printf("La cantidad de letras de la palabra \"%v\" es: %d\n", palabra, len(palabra))
	for _, v := range palabra {
		fmt.Printf("%c\n", v)
	}

	//Ejercicio 2 - Préstamo
	//Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene
	//ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea
	//mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos
	//que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
	//Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
	//Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
	var (
		edadCliente       = 25
		estaEmpleado      = true
		aniosDeAntiguedad = 1.01
		sueldo            = 150000.0
	)
	if !(edadCliente > 22 && estaEmpleado && aniosDeAntiguedad > 1) {
		fmt.Println("No pudiste acceder al prestamo")
	} else {
		fmt.Println("Pudiste acceder al prestamo")
		if sueldo > 100000 {
			fmt.Println("Y no se te cobrara interes")
		} else {
			fmt.Println("Y se te cobrara interes")
		}
	}

	// Ejercicio 3 - A qué mes corresponde
	//Realizar una aplicación que reciba  una variable con el número del mes.
	//Según el número, imprimir el mes que corresponda en texto.
	//¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	//Ej: 7, Julio.
	//Nota: Validar que el número del mes, sea correcto.
	mapMeses := make(map[int]string)
	nombreMeses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	numeroMesABuscar := 1
	for i := range nombreMeses {
		mapMeses[i+1] = nombreMeses[i]
	}
	if numeroMesABuscar > 0 && numeroMesABuscar <= 12 {
		println(numeroMesABuscar, mapMeses[numeroMesABuscar])
	} else {
		println("El numero de mes ingresado no existe")
	}

	// Ejercicio 4 - Qué edad tiene...
	//Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.
	//Por otro lado también es necesario:
	//Saber cuántos de sus empleados son mayores de 21 años.
	//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	//Eliminar a Pedro del mapa.
	var (
		employees       = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
		contadorMayores = 0
	)
	fmt.Printf("La edad de Benjamin es: %v\n", employees["Benjamin"])
	for _, v := range employees {
		if v > 21 {
			contadorMayores++
		}
	}
	fmt.Printf("La cantidad de empleados mayores a 21 es: %d", contadorMayores)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}
