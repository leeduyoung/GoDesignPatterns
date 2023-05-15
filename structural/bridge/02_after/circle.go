package main

type Circle struct {
	x, y, radius int
	drawingAPI   DrawingAPI
}

func NewCircle(x, y, radius int, api DrawingAPI) *Circle {
	return &Circle{
		x:          x,
		y:          y,
		radius:     radius,
		drawingAPI: api,
	}
}

func (c Circle) Draw() {
	c.drawingAPI.DrawCircle(c.x, c.y, c.radius)
}
