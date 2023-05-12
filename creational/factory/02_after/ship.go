package main

type Ship struct {
	name  string
	color string
	logo  string
}

func NewShip(name, color, logo string) *Ship {
	return &Ship{
		name:  name,
		color: color,
		logo:  logo,
	}
}

func (s *Ship) Name() string {
	return s.name
}

func (s *Ship) SetName(name string) {
	s.name = name
}

func (s *Ship) Color() string {
	return s.color
}

func (s *Ship) SetColor(color string) {
	s.color = color
}

func (s *Ship) Logo() string {
	return s.name
}

func (s *Ship) SetLogo(logo string) {
	s.logo = logo
}
