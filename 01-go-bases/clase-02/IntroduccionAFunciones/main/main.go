package main

import "errors"

const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
)

func main() {
	println(operacionAritmetica(6, 2, "+"))
	println(operacionAritmetica(6, 2, "-"))
	println(operacionAritmetica(6, 2, "*"))
	println(operacionAritmetica(6, 2, "/"))
	s, r, m, d := operacionAritmeticaMultipleRetorno(6, 2)
	println("Suma:\t", s)
	println("Resta:\t", r)
	println("Multip:\t", m)
	println("Divis:\t", d)
	sNombrada, rNombrada, mNombrada, dNombrada := retornoValoresNombrados(6, 2)
	println("Suma:\t", sNombrada)
	println("Resta:\t", rNombrada)
	println("Multip:\t", mNombrada)
	println("Divis:\t", dNombrada)
	result, err := division(6, 0)
	if err != nil {
		println(err.Error())
	} else {
		println(result)
	}
}

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Suma:
		return valor1 + valor2
	case Resta:
		return valor1 - valor2
	case Multip:
		return valor1 * valor2
	case Divis:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
}

func operacionAritmeticaMultipleRetorno(valor1, valor2 float64) (float64, float64, float64, float64) {
	suma := valor1 + valor2
	resta := valor1 - valor2
	multip := valor1 * valor2
	var divis float64
	if valor2 != 0 {
		divis = valor1 / valor2
	}
	return suma, resta, multip, divis
}

func division(valor1, valor2 float64) (float64, error) {
	if valor2 == 0 {
		return 0, errors.New("No es posible dividir en 0 ")
	}
	return valor1 / valor2, nil
}

func retornoValoresNombrados(valor1, valor2 float64) (suma float64, resta float64, multip float64, divis float64) {
	suma = valor1 + valor2
	resta = valor1 - valor2
	multip = valor1 * valor2
	if valor2 != 0 {
		divis = valor1 / valor2
	}
	return
}

// Trying if I can send many multiple parameters to a single funcion
func test(valor1 float64, operador ...any) float64 {
	return 0
}

//It cant receive many multiple parameters, but I can give It a parameter with any type, so Is nearly the same
