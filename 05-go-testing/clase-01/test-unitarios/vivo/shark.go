package hunt

import "fmt"

var (
	ErrNoPrey            = fmt.Errorf("there is no prey")
	ErrTired             = fmt.Errorf("cannot hunt, i am really tired")
	ErrNotHungry         = fmt.Errorf("cannot hunt, i am not hungry")
	ErrCouldNotCatchPrey = fmt.Errorf("could not catch it")
)

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func (s *Shark) Hunt(p *Prey) error {
	if p == nil {
		return ErrNoPrey
	}
	if s.tired {
		return ErrTired
	}
	if !s.hungry {
		return ErrNotHungry
	}
	if p.speed >= s.speed {
		s.tired = true
		return ErrCouldNotCatchPrey
	}

	s.hungry = false
	s.tired = true
	return nil
}
