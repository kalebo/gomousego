package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	modePtr := flag.String("mode", "random", "Mode to use.")
	flag.Parse()

	var strategy IStrategy

	switch *modePtr {
	case "random":
		fmt.Printf("Unimplemented mode: %s\n", *modePtr)
	case "schizo":
		fmt.Printf("Unimplemented mode: %s\n", *modePtr)
	case "momentum":
		// Cole's original algo
		fmt.Println("Running in mode Momentum")
		strategy = NewMomentumStrategy(10)
	}

	runStrategy(strategy, 10)

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

func NewMomentumStrategy(damping int) *MomentumStrategy {
	x, y := robotgo.GetMousePos()
	return &MomentumStrategy{Pair{x, y}, Pair{0, 0}, damping}
}
