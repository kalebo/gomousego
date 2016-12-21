package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type ScrollStrategy struct {
	duration int
}

func NewScrollStrategy() *ScrollStrategy {
	return &ScrollStrategy{3000}
}

func (s *ScrollStrategy) Step() {
	r := rand.Intn(100)
	amount := rand.Intn(5)

	if r < 15 {
		robotgo.ScrollMouse(amount, "up")
	} else if r < 30 {
		robotgo.ScrollMouse(amount, "down")
	}
}

func (s *ScrollStrategy) StepDuration() time.Duration {
	return time.Duration(s.duration) * time.Millisecond
}
