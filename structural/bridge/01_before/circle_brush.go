package main

import "fmt"

type CircleBrush struct {
	x, y, radius int
}

func NewCircleBrush(x, y, radius int) *CircleBrush {
	return &CircleBrush{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (cb CircleBrush) Draw() {
	fmt.Println("drawing brush circle...")
}
