package main

import "fmt"

type SquarePencil struct {
	x, y, radius int
}

func NewSquarePencil(x, y, radius int) *SquarePencil {
	return &SquarePencil{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (sp SquarePencil) Draw() {
	fmt.Println("drawing pencil square...")
}
