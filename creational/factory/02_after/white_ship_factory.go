package main

type WhiteShipFactory struct {
}

func NewWhiteShipFactory() *WhiteShipFactory {
	return &WhiteShipFactory{}
}

func (wsf *WhiteShipFactory) CreateShip(name string) *Ship {
	// Customizing for specific name
	var (
		color = "white"
		logo  = "whiteShipLogo"
	)

	// Create ship
	return NewShip(name, color, logo)
}
