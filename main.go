package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	modePtr := flag.String("mode", "momentum", "Mode to use.")
	flag.Parse()

	var strategy IStrategy

	switch *modePtr {
	case "random":
		// TODO
	case "schizo":
		// TODO
	case "momentum":
		// Cole's original algo
		strategy = NewMomentumStrategy(15)
	}

	if strategy != nil {
		fmt.Printf("Running in mode: %s\n", *modePtr)
		runStrategy(strategy, 10)
	} else {
		fmt.Printf("Unimplemented mode: %s\n", *modePtr)
	}

}

func runStrategy(s IStrategy, interval int) {
	t := time.NewTicker(time.Millisecond * time.Duration(interval))
	for {
		s.Run()
		<-t.C
	}
}

type IStrategy interface {
	Run()
}

type Pair struct {
	x, y int
}

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
