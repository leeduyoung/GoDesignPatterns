package main

import "fmt"

type DrawingPencilAPI struct {
}

func NewDrawingPencilAPI() *DrawingPencilAPI {
	return &DrawingPencilAPI{}
}

func (d DrawingPencilAPI) DrawCircle(x, y, radius int) {
	fmt.Println("drawing pencil circle...")
}

func (d DrawingPencilAPI) DrawSquare(x, y, side int) {
	fmt.Println("drawing pencil square...")
}
