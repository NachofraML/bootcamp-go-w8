package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

// Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
// Para ello, cuentan con todo el detalle necesario en un archivo .txt.
// Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo,
// no han pasado el archivo a leer por nuestro programa.
// Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt” (recordá lo visto sobre el
// pkg “os”).
// Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso,
// el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando
// el mensaje “el archivo indicado no fue encontrado o está dañado”.
// Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

////Ejercicio 2 - Imprimir datos
////A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
////Ahora que el archivo sí existe, el panic no debe ser lanzado.
////Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
////Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
////En el caso de no poder leerlo, se debe lanzar un “panic”.
////Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer”
////que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.

//Ejercicio 3 - Registrando clientes
//El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
//Los datos requeridos son:
//Legajo
//Nombre
//DNI
//Número de teléfono
//Domicilio
//Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los
//datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí.
//Ese error deberá:
//1.- generar un panic;
//2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la
//ejecución del programa normalmente.
//Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que
//todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos,
//dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero
//(recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
//Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes
//mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir
//con este requerimiento.

//Requerimientos generales:
//Utilizá recover para recuperar el valor de los panics que puedan surgir
//Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
//Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la
//validación pertinente para el caso de error retornado).

type Client struct {
	Legajo    string
	Nombre    string
	DNI       int
	Telefono  int
	Domicilio string
}

var (
	FilePath = "./ex/customers.txt"
	Clients  = &[]Client{}
)

// Imaginemos que hice un error, es facil pero no lo veo necesario para practicar panics

func main() {
	defer func() {
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}()
	var err error

	// Segundo y tercer ejercicio
	client1 := Client{
		Legajo:    "Legajo 1",
		Nombre:    "Nacho",
		DNI:       12345678,
		Telefono:  12345678,
		Domicilio: "Hi 1",
	}
	client2 := Client{
		Legajo:    "Legajo 2",
		Nombre:    "Nacha",
		DNI:       12345678,
		Telefono:  23456789,
		Domicilio: "Hi 2",
	}
	_, err = saveCient(client1)
	if err != nil {
		fmt.Println(err)
	}

	_, err = saveCient(client2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Clients)

	// Primer ejercicio
	err = readFile()
	if err != nil {
		err = createFile()
		if err != nil {
			panic(err)
		}
	}

	err = writeClientsToFile(*Clients)
	if err != nil {
		panic(err)
	}
}

func readFile() (err error) {
	_, err = os.Open(FilePath)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	}

	return
}

func createFile() (err error) {
	_, err = os.Create(FilePath)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	}
	return
}

func writeClientsToFile(clients []Client) (err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Create(FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	if err = encoder.Encode(clients); err != nil {
		panic(err)
	}

	_, err = io.Copy(file, buf)
	if err != nil {
		panic(err)
	}

	return
}

func saveCient(client Client) (ok bool, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	ok, err = validateClientData(client)
	if err != nil {
		panic(err)
	}

	for _, c := range *Clients {
		if c.DNI == client.DNI {
			panic("Error: el cliente ya existe")
		}
	}

	*Clients = append(*Clients, client)
	return
}

func validateClientData(client Client) (ok bool, err error) {
	clientValue := reflect.ValueOf(client)
	for i := 0; i < clientValue.NumField(); i++ {
		fieldValue := clientValue.Field(i)
		if fieldValue.IsZero() {
			panic(fmt.Errorf("la variable %v es nula", fieldValue))
		}
	}
	ok = true
	return
}
