package main

type Ship struct {
	name   string
	color  string
	logo   string
	wheel  Wheel
	anchor Anchor
}

func NewShip(name, color, logo string, wheel Wheel, anchor Anchor) *Ship {
	return &Ship{
		name:   name,
		color:  color,
		logo:   logo,
		wheel:  wheel,
		anchor: anchor,
	}
}

func (s *Ship) Name() string {
	return s.name
}

func (s *Ship) Color() string {
	return s.color
}

func (s *Ship) Logo() string {
	return s.name
}

func (s *Ship) Wheel() Wheel {
	return s.wheel
}

func (s *Ship) Anchor() Anchor {
	return s.anchor
}
