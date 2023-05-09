package shark

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-dobles/vivo/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}
