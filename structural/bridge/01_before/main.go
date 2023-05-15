package main

func main() {
	// Create a Circle with DrawingPencilAPI
	circleWithPencil := NewCirclePencil(2, 3, 5)
	circleWithPencil.Draw()

	// Create a Square with DrawingPencilAPI
	squareWithPencil := NewSquarePencil(4, 1, 7)
	squareWithPencil.Draw()

	// Create a Square with DrawingBrushAPI
	squareWithBrush := NewSquareBrush(4, 1, 7)
	squareWithBrush.Draw()
}
