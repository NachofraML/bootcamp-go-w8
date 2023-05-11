package shark

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/vivo/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}
