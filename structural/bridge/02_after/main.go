package main

func main() {
	// Create a Circle with DrawingPencilAPI
	circleWithPencil := NewCircle(2, 3, 5, NewDrawingPencilAPI())
	circleWithPencil.Draw()

	// Create a Square with DrawingPencilAPI
	squareWithPencil := NewSquare(4, 1, 7, NewDrawingPencilAPI())
	squareWithPencil.Draw()

	// Create a Square with DrawingBrushAPI
	squareWithBrush := NewSquare(4, 1, 7, NewDrawingBrushAPI())
	squareWithBrush.Draw()
}
