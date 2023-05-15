package main

type DrawingAPI interface {
	DrawCircle(x, y, radius int)
	DrawSquare(x, y, side int)
}
