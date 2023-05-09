package simulator

import (
	"fmt"
)

// Ejercicio 3
// Crear un mock para el simulator. El mock debe implementar simular la implementación del método CanCatch
// y un spy del método GetLinearDistance.

// CatchSimulatorMock I'm using Args for studying reasons, in this case is better something more structured, because we don't have many
// arguments
type CatchSimulatorMock struct {
	CanCatchMock          Args
	GetLinearDistanceMock Args
}

func (r *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	timeToCatch := r.CanCatchMock.Float64(0)
	maxTimeToCatch := r.CanCatchMock.Float64(1)
	fmt.Printf("Time to catch: %.2f seconds\n", timeToCatch)
	return timeToCatch > 0 && timeToCatch <= maxTimeToCatch
}

func (r *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	res := r.GetLinearDistanceMock.Float64(0)
	fmt.Printf("Distance: %.2f meters\n", res)
	return res
}

func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

// NewArgs ------------------------------------------------------------
func NewArgs(a ...any) Args {
	return Args{a}
}

// Args is a slice of arguments of type any.
type Args struct {
	a []any
}

// Get returns the argument at the given index.
func (a *Args) Get(i int) any {
	return a.a[i]
}
func (a *Args) String(i int) string {
	return a.a[i].(string)
}
func (a *Args) Int(i int) int {
	return a.a[i].(int)
}
func (a *Args) Float64(i int) float64 {
	return a.a[i].(float64)
}
func (a *Args) Bool(i int) bool {
	return a.a[i].(bool)
}

// pointer cases
func (a *Args) Error(i int) error {
	var e error
	if a.a[i] != nil {
		e = a.a[i].(error)
	}
	return e
}
