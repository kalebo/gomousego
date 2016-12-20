package main

import (
	"flag"
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	modePtr := flag.String("mode", "random", "Mode to use.")
	flag.Parse()

	fmt.Printf("Using mode: %s\n", *modePtr)

	robotgo.ScrollMouse(10, "up")
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
