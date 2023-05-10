package prey

// Ejercicio 2
// Crear stubs de prey para poder realizar los tests. Se deben poder cubrir todos los casos del método GetSpeed.

type TunaStub struct {
	GetSpeedStub func() float64
}

func (p *TunaStub) GetSpeed() float64 {
	return p.GetSpeedStub()
}

func CreateTunaStub() *TunaStub {
	return &TunaStub{}
}
