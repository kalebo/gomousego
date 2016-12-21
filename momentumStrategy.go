package main

import "github.com/go-vgo/robotgo"

type MomentumStrategy struct {
	previous Pair
	momentum Pair
	damping  int
}

func NewMomentumStrategy(damping int) *MomentumStrategy {
	x, y := robotgo.GetMousePos()
	return &MomentumStrategy{Pair{x, y}, Pair{0, 0}, damping}
}

func (s *MomentumStrategy) Run() {
	x, y := robotgo.GetMousePos()

	dif := Pair{
		s.previous.x - x,
		s.previous.y - y,
	}

	s.momentum = Pair{
		s.momentum.x - (dif.x / s.damping),
		s.momentum.y - (dif.y / s.damping),
	}

	robotgo.MoveMouse(x+s.momentum.x, y+s.momentum.y)
	s.previous = Pair{x, y}
}
