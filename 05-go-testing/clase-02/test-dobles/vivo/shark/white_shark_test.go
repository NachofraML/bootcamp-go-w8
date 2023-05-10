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

// Test with args
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
		simulatorMock.CanCatchMock = simulator.NewArgs(true)
		// Changing linear distance does not affect this test, but is necessary to be initialized
		simulatorMock.GetLinearDistanceMock = simulator.NewArgs(0.0)

		whiteShark := CreateWhiteShark(simulatorMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Failed Hunt", func(t *testing.T) {
		// Arrange
		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorMock()
		// inputting as Args timeToCatch & maxTimeToCatch
		simulatorMock.CanCatchMock = simulator.NewArgs(false)
		// Changing linear distance does not affect this test, but is necessary to be initialized
		simulatorMock.GetLinearDistanceMock = simulator.NewArgs(0.0)

		whiteShark := CreateWhiteShark(simulatorMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.ErrorIs(t, err, ErrCouldNotHuntPrey)
	})
}

// Test with Testify mock
func TestBetterWhiteSharkHunt(t *testing.T) {
	// TODO: Check if GetLinearDistance was called
	t.Run("Successful Hunt", func(t *testing.T) {
		// Arrange
		var distance, speed, catchSpeed float64

		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorBetterMock()
		// Mocked MaxTimeToCatch
		simulatorMock.MaxTimeToCatch = 15
		// Mocked distance response
		distance = 14.4
		// Default speed of shark
		speed = 144
		// Prey speed
		catchSpeed = tunaStub.GetSpeedStub()
		// 	Defining happy path distance, speed & catchSpeed CanCatch response
		simulatorMock.
			On("CanCatch", distance, speed, catchSpeed, simulatorMock.MaxTimeToCatch).
			Return(true)

		// GetLinearDistance is typically used to calculate the linear distance between the hunter's current position
		// and the target position during the hunting process.
		// However, in this section of the code, we're using the shark's predefined position (initialized to zero)
		// to compare it with our test position (also initialized to zero), so we can finally return the mocked distance
		// because they are always equal.
		position := [2]float64{0, 0}
		simulatorMock.
			On("GetLinearDistance", position).
			Return(distance)

		whiteShark := whiteShark{speed, position, simulatorMock}

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.NoError(t, err)
		// -> expected calls
		simulatorMock.AssertExpectations(t)
	})

	t.Run("Failed Hunt By Speed", func(t *testing.T) {
		// Arrange
		var distance, speed, catchSpeed float64

		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorBetterMock()
		// Mocked MaxTimeToCatch
		simulatorMock.MaxTimeToCatch = 15
		// Mocked distance response
		distance = 14.4
		// Default speed of shark
		speed = 9
		// Prey speed
		catchSpeed = tunaStub.GetSpeedStub()
		// 	Defining happy path distance, speed & catchSpeed CanCatch response
		simulatorMock.
			On("CanCatch", distance, speed, catchSpeed, simulatorMock.MaxTimeToCatch).
			Return(false)

		// GetLinearDistance is typically used to calculate the linear distance between the hunter's current position
		// and the target position during the hunting process.
		// However, in this section of the code, we're using the shark's predefined position (initialized to zero)
		// to compare it with our test position (also initialized to zero), so we can finally return the mocked distance
		// because they are always equal.
		position := [2]float64{0, 0}
		simulatorMock.
			On("GetLinearDistance", position).
			Return(distance)

		whiteShark := whiteShark{speed, position, simulatorMock}

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, ErrCouldNotHuntPrey.Error())
		// -> expected calls
		simulatorMock.AssertExpectations(t)
	})

	t.Run("Failed Hunt By Distance", func(t *testing.T) {
		// Arrange
		var distance, speed, catchSpeed float64

		tunaStub := prey.CreateTunaStub()
		tunaStub.GetSpeedStub = func() (speed float64) {
			speed = 10.0
			return
		}

		simulatorMock := simulator.NewCatchSimulatorBetterMock()
		// Mocked MaxTimeToCatch
		simulatorMock.MaxTimeToCatch = 15
		// Mocked distance response
		distance = 20
		// Default speed of shark
		speed = 11
		// Prey speed
		catchSpeed = tunaStub.GetSpeedStub()
		// 	Defining happy path distance, speed & catchSpeed CanCatch response
		simulatorMock.
			On("CanCatch", distance, speed, catchSpeed, simulatorMock.MaxTimeToCatch).
			Return(false)
		// GetLinearDistance is typically used to calculate the linear distance between the hunter's current position
		// and the target position during the hunting process.
		// However, in this section of the code, we're using the shark's predefined position (initialized to zero)
		// to compare it with our test position (also initialized to zero), so we can finally return the mocked distance
		// because they are always equal.
		position := [2]float64{0, 0}
		simulatorMock.
			On("GetLinearDistance", position).
			Return(distance)

		whiteShark := whiteShark{speed, position, simulatorMock}

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, ErrCouldNotHuntPrey.Error())
		// -> expected calls
		simulatorMock.AssertExpectations(t)
	})
}
