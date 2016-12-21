package main

import (
	"flag"
	"fmt"
)

type IStrategy interface {
	Run()
}

type Pair struct {
	x, y int
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
		// Cole's original algo
		strategy = NewMomentumStrategy(15)
	case "scroll":
		strategy = NewScrollStrategy()
	}

	if strategy != nil {
		fmt.Printf("Running in mode: %s\n", *modePtr)
		strategy.Run()
	} else {
		fmt.Printf("Unimplemented mode: %s\n", *modePtr)
	}

}
