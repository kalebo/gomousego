package main

import (
	"flag"
	"fmt"
	"time"
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
