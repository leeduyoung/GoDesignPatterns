package main

import "fmt"

type SquareBrush struct {
	x, y, radius int
}

func NewSquareBrush(x, y, radius int) *SquareBrush {
	return &SquareBrush{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (sb SquareBrush) Draw() {
	fmt.Println("drawing brush square...")
}
