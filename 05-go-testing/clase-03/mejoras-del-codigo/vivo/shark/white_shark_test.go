package shark

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/vivo/pkg/storage"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/vivo/simulator"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Manual Mock Test
func TestWhiteSharkHuntIntegration(t *testing.T) {
	t.Run("Successful Hunt", func(t *testing.T) {
		// Arrange
		storageMockConfig := map[string]interface{}{
			"white_shark_speed": 144.0,
			"white_shark_x":     10.0,
			"white_shark_y":     10.0,
			"tuna_speed":        10.0,
		}
		storageMock := storage.NewStorageMock()
		storageMock.Values = storageMockConfig

		tunaStub := prey.CreateTuna(storageMock)

		maxTimeToCatch := 30.0
		simulatorMock := simulator.NewCatchSimulator(maxTimeToCatch)

		whiteShark := CreateWhiteShark(simulatorMock, storageMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("Failed Hunt By Speed", func(t *testing.T) {
		// Arrange
		storageMockConfig := map[string]interface{}{
			"white_shark_speed": 9.0,
			"white_shark_x":     10.0,
			"white_shark_y":     10.0,
			"tuna_speed":        10.0,
		}
		storageMock := storage.NewStorageMock()
		storageMock.Values = storageMockConfig

		tunaStub := prey.CreateTuna(storageMock)

		maxTimeToCatch := 30.0
		simulatorMock := simulator.NewCatchSimulator(maxTimeToCatch)

		whiteShark := CreateWhiteShark(simulatorMock, storageMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, ErrCouldNotHuntPrey.Error())
	})
	t.Run("Failed Hunt By Distance", func(t *testing.T) {
		// Arrange
		storageMockConfig := map[string]interface{}{
			"white_shark_speed": 11.0,
			"white_shark_x":     0.0,
			"white_shark_y":     350.0,
			"tuna_speed":        10.0,
		}
		storageMock := storage.NewStorageMock()
		storageMock.Values = storageMockConfig

		tunaStub := prey.CreateTuna(storageMock)

		maxTimeToCatch := 30.0
		simulatorMock := simulator.NewCatchSimulator(maxTimeToCatch)

		whiteShark := CreateWhiteShark(simulatorMock, storageMock)

		// Act
		err := whiteShark.Hunt(tunaStub)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, ErrCouldNotHuntPrey.Error())
	})
}

// TODO: Testify Mock Test
