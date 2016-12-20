package main

import (
	"flag"
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	modePtr := flag.String("mode", "random", "Mode to use.")
	flag.Parse()

	fmt.Fprintln("Using mode: %s", *modePtr)

	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
