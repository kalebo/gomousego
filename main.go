package main

import (
	"flag"
	"fmt"
	"time"
)

type IStrategy interface {
	Step()
	StepDuration() time.Duration
}

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
		// Cole's original daemon mouse
		strategy = NewMomentumStrategy(15)
	case "scroll":
		strategy = NewScrollStrategy()
	}

	if strategy != nil {
		fmt.Printf("Running in mode: %s\n", *modePtr)
		RunStrategy(strategy)
	} else {
		fmt.Printf("Unimplemented mode: %s\n", *modePtr)
	}
}

func RunStrategy(s IStrategy) {
	t := time.NewTicker(s.StepDuration())
	for {
		s.Step()
		<-t.C
	}
}
