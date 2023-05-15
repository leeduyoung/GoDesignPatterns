package main

import "fmt"

type CirclePencil struct {
	x, y, radius int
}

func NewCirclePencil(x, y, radius int) *CirclePencil {
	return &CirclePencil{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (cp CirclePencil) Draw() {
	fmt.Println("drawing pencil circle...")
}
