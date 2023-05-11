package shark

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) (error, float64)
	Configure(position [2]float64, speed float64)
}
