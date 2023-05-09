package shark

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-dobles/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-dobles/vivo/simulator"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Ejercicio 1
// Si se fijan bien en el código, ya existe un test para el struct whiteShark con el que trabajaremos. Sin embargo,
// el test no está bien realizado. ¿Por qué?
// Borrar el test antes de continuar con el resto de los ejercicios.

// Ejercicio 4
// Realizar test unitarios del método Hunt del tiburón blanco, cubriendo todos los casos posibles, usando los
// stubs y mocks creados anteriormente:
// El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta. Hacer un assert de que el
// método GetLinearDistance fue llamado.
// El tiburón no logra cazar el atún al ser más lento.
// El tiburón no logra cazar el atún por estar a una distancia muy larga, a pesar de ser más veloz.

func TestWhiteSharkHunt(t *testing.T) {
	t.Run("Successful Hunt", func(t *testing.T) {
		// Arrange
		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorMock()
		// inputting as Args timeToCatch & maxTimeToCatch
		simulatorMock.CanCatchMock = simulator.NewArgs(10.0, 15.0)
		// inputting as Args distance in meters
		simulatorMock.GetLinearDistanceMock = simulator.NewArgs(10.0)

		whiteShark := CreateWhiteShark(simulatorMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Failed Hunt By Speed", func(t *testing.T) {
		// Arrange
		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorMock()
		// inputting as Args timeToCatch & maxTimeToCatch
		simulatorMock.CanCatchMock = simulator.NewArgs(20.0, 15.0)
		// inputting as Args distance in meters
		simulatorMock.GetLinearDistanceMock = simulator.NewArgs(10.0)

		whiteShark := CreateWhiteShark(simulatorMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.ErrorIs(t, err, ErrCouldNotHuntPrey)
	})
	t.Run("Failed Hunt By Distance", func(t *testing.T) {
	})
}
