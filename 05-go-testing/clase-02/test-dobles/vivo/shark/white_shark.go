package shark

import (
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-dobles/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-dobles/vivo/simulator"
	"math/rand"
	"time"
)

var (
	ErrCouldNotHuntPrey = fmt.Errorf("could not hunt the prey")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type whiteShark struct {
	// speed in m/s
	speed float64
	// position of the shark in the map of 500 * 500 meters
	position [2]float64
	// simulator
	simulator simulator.CatchSimulator
}

func (w *whiteShark) Hunt(prey prey.Prey) error {
	if w.simulator.CanCatch(w.simulator.GetLinearDistance(w.position), w.speed, prey.GetSpeed()) {
		fmt.Println("ñam ñam")
		return nil
	}
	return ErrCouldNotHuntPrey
}

func CreateWhiteShark(simulator simulator.CatchSimulator) Shark {
	return &whiteShark{
		speed:     144,
		position:  getInitialPosition(),
		simulator: simulator,
	}
}

func getInitialPosition() [2]float64 {
	x := rand.Float64() * 500
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64() * 500
	return [2]float64{x, y}
}
