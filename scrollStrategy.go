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
	rint := rand.Intn(100)

	if rint < 25 {
		robotgo.ScrollMouse(10, "up")
	} else if rint < 50 {
		robotgo.ScrollMouse(10, "down")
	}
}
