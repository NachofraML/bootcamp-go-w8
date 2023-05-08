package hunt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test By Method (I could do subtests but is similar)
func TestSharkHuntsSuccessfully(t *testing.T) {
	//Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  11,
	}
	prey := Prey{
		name:  "Nemo",
		speed: 10,
	}

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.NoError(t, err)
	assert.False(t, shark.hungry)
	assert.True(t, shark.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//Arrange
	shark := Shark{
		hungry: false,
		tired:  true,
		speed:  100,
	}
	prey := Prey{
		name:  "Nemo",
		speed: 10,
	}

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.Error(t, err)
	assert.ErrorContains(t, err, "cannot hunt, i am really tired")
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//Arrange
	shark := Shark{
		hungry: false,
		tired:  false,
		speed:  100,
	}
	prey := Prey{
		name:  "Nemo",
		speed: 10,
	}

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.Error(t, err)
	assert.ErrorContains(t, err, "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	prey := Prey{
		name:  "Nemo",
		speed: 11,
	}

	//Act
	err := shark.Hunt(&prey)

	//Assert
	assert.Error(t, err)
	assert.ErrorContains(t, err, "could not catch it")
}

func TestSharkHuntNilPrey(t *testing.T) {
	//Arrange
	shark := Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	var prey *Prey

	//Act
	err := shark.Hunt(prey)

	//Assert
	assert.Error(t, err)
	assert.ErrorContains(t, err, "there is no prey")
}

// Tabla Driven Test Example
func TestSharkHuntTDT(t *testing.T) {
	type input struct {
		shark Shark
		prey  *Prey
	}
	type output struct {
		err error
	}
	//Arrange
	testCases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name:   "TestSharkCannotHuntBecauseIsTired",
			input:  input{shark: Shark{hungry: true, tired: true, speed: 10}, prey: &Prey{name: "Nemo", speed: 10}},
			output: output{err: ErrTired},
		},
		{
			name:   "TestSharkCannotHuntBecaisIsNotHungry",
			input:  input{shark: Shark{hungry: false, tired: false, speed: 10}, prey: &Prey{name: "Nemo", speed: 10}},
			output: output{err: ErrNotHungry},
		},
		{
			name:   "TestSharkCannotReachThePrey",
			input:  input{shark: Shark{hungry: true, tired: false, speed: 10}, prey: &Prey{name: "Nemo", speed: 11}},
			output: output{err: ErrCouldNotCatchPrey},
		},
		{
			name:   "TestSharkHuntNilPrey",
			input:  input{shark: Shark{hungry: true, tired: true, speed: 10}, prey: nil},
			output: output{err: ErrNoPrey},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			err := tc.input.shark.Hunt(tc.input.prey)

			// Assert
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.output.err)
		})
	}
}
