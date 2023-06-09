package main

import "fmt"

//Ejercicio 1 - Imprimí tu nombre
//Tendrás que:
//Crear una aplicación donde tengas como variable tu nombre y dirección.
//Imprimir en consola el valor de cada una de las variables.

//Ejercicio 2 - Clima
//Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
//Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
//Imprimí los valores de las variables en consola.
//¿Qué tipo de dato le asignarías a las variables?

//Ejercicio 3 - Declaración de variables
//Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder
//brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
//Necesita ayuda para:
//Detectar cuáles de estas variables que declaró el alumno son correctas.
//Corregir las incorrectas.
//  var 1nombre string
//  var apellido string
//  var int edad
//  1apellido := 6
//  var licencia_de_conducir = true
//  var estatura de la persona int
//  cantidadDeHijos := 2

// Ejercicio 4 - Tipos de datos
//Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go,
//pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un
//desarrollador experimentado que pueda:
//Verificar su código y realizar las correcciones necesarias.
//  var apellido string = "Gomez"
//  var edad int = "35"
//  boolean := "false";
//  var sueldo string = 45857.90
//  var nombre string = "Julián"

var (
	temperatura float64 = 30.7
	humedad     float64 = 20.0
	presion     float64 = 0.0
	//  var 1nombre string MAL, NO PUEDE EMPEZAR CON NUMERO
	//  var apellido string BIEN
	//  var int edad MAL, INT ES RESERVADA Y EDAD NO ES UN TIPO
	//  1apellido := 6 MAL, NO PUEDE EMPEZAR CON NUMERO
	//  var licencia_de_conducir = true BIEN
	//  var estatura de la persona int MAL, NO PUEDE TENER ESPACIOS
	//  cantidadDeHijos := 2 BIEN
	//  var apellido string = "Gomez" BIEN
	//  var edad int = "35" MAL, ES UN INT NO STRING
	//  boolean := "false"; BIEN PERO NO ESTA USANDO TIPO BOOL
	//  var sueldo string = 45857.90 MAL, ES UN FLOAT
	//  var nombre string = "Julián" BIEN
)

func main() {
	nombre, direccion := "Ignacio F Mosca", "Direccion 1234"
	fmt.Println(nombre, " ", direccion)
	fmt.Printf("Temp: %v, Humedad: %v, Presion: %v", temperatura, humedad, presion)
}
