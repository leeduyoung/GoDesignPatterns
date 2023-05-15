package main

import "fmt"

type DrawingBrushAPI struct {
}

func NewDrawingBrushAPI() *DrawingBrushAPI {
	return &DrawingBrushAPI{}
}

func (d DrawingBrushAPI) DrawCircle(x, y, radius int) {
	fmt.Println("drawing brush circle...")
}

func (d DrawingBrushAPI) DrawSquare(x, y, side int) {
	fmt.Println("drawing brush square...")
}
