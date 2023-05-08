package prices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Como cumplir con cada aspecto del Principio FIRST

// FAST: test rapidos, diseñar test unitarios cada test debe cubrir solo la unidad minima de codigo, esto permite que cada test se ejecute de forma rápida
func TestCalcPrice(t *testing.T) {
	//probamos solo la funcion calPrice y le pasamos un articulo mockeado
	artMock := Article{
		Name:      "dummy",
		CostPrice: 1,
		Tax:       1,
	}
	precioEsperado := float64(artMock.CostPrice + artMock.Tax)
	price, err := calcPrice(artMock)
	// se valida que el error sea nulo y que el precio sea el valor esperado
	assert.Nil(t, err)
	assert.Equal(t, precioEsperado, price)
}

// INDEPENDENT: Cada test es indepentiente del otro, podriamos ejecutar cualquier otro test sin depender de ningun test anterio
func TestGetArticle(t *testing.T) {
	_, err := getArticle("ASX123")
	// se valida que el error sea nulo y que el precio sea el valor esperado
	// este test no depende de ningun otro, y se pueden ejecutar en cualquier orden
	assert.Nil(t, err)
}

// REPEATABLE: Estas pruebas pueden ser repetida en cualquier otro server y deben arrojar el mismo resultado.
// si en vez de tener una funcion que emula la obtencion de un articulo, tuvieramos un DB, habria que crear un mock de la misma para garantizar que el test es Repetible
func TestGetArticleError(t *testing.T) {
	// en esete test enviamos un ID inexistente validando que retorne lo esperado
	_, err := getArticle("QWEQWE")
	// se valida que el error sea nulo y que el precio sea el valor esperado
	// este test es repetible en cualquier escenario
	assert.Equal(t, errors.New("article not found"), err)
}

// SELF-VALIDATING (auto evaluable). Todos los tests se autoevaluan usando assert, existen otros packages que permiten hacer las auto validaciones (require, por ejm)
// ejecutar todas las validaciones posibles, garantizar probar los distintos flujos del programa
// En el siguiente test vamos a testear el Error de la funcion calcPrice cuando un articulo no tiene impuesto definido
func TestCalcPriceError(t *testing.T) {
	//probamos solo la funcion calPrice y le pasamos un articulo mockeado
	artMock := Article{
		Name:      "dummy",
		CostPrice: 1,
		Tax:       0,
	}
	precioEsperado := float64(0)
	price, err := calcPrice(artMock)
	// se valida que el error sea distinto de nulo y que el precio sea 0
	assert.NotNil(t, err)
	assert.Equal(t, precioEsperado, price)
}

// TIMELY (oportuno)Es conveniente desarrollar primero los tests unitarios antes de de desarrollar el codigo del producto
// Un test bien hecho practicamente determina como debe ser el código del producto.
