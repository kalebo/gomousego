package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type ScrollStrategy struct {
}

func NewScrollStrategy() *ScrollStrategy {
	return &ScrollStrategy{}
}

func (s *ScrollStrategy) Run() {
	t := time.NewTicker(time.Millisecond * time.Duration(3000))
	for {
		s.Step()
		<-t.C
	}
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
