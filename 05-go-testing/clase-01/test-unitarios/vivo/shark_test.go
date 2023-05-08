package hunt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.ErrorContains(t, err, "prey does not exist")
}

// Test if we cannot change the code, so it does not panic
//func TestSharkHuntNilPrey(t *testing.T) {
//	//Arrange
//	shark := Shark{
//		hungry: true,
//		tired:  false,
//		speed:  10,
//	}
//	var prey *Prey
//
//	defer func() {
//		if r := recover(); r != nil {
//			// Here we can handle the panic like we want
//			require.NotNil(t, prey)
//		}
//	}()
//
//	//Act
//	err := shark.Hunt(prey)
//
//	//Assert
//	assert.Error(t, err)
//}
