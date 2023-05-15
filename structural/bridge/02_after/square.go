package main

type Square struct {
	x, y, side int
	drawingAPI DrawingAPI
}

func NewSquare(x, y, side int, api DrawingAPI) *Square {
	return &Square{
		x:          x,
		y:          y,
		side:       side,
		drawingAPI: api,
	}
}

func (s Square) Draw() {
	s.drawingAPI.DrawSquare(s.x, s.y, s.side)
}
